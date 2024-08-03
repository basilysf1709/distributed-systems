const std = @import("std");

const VirtualNode = struct {
    id: []u8,
    position: u64,
    physical_node: []u8,
};

const ConsistentHash = struct {
    nodes: std.ArrayList(VirtualNode),
    allocator: std.mem.Allocator,
    virtual_nodes_per_server: u32,

    pub fn init(allocator: std.mem.Allocator, virtual_nodes: u32) ConsistentHash {
        return ConsistentHash{
            .nodes = std.ArrayList(VirtualNode).init(allocator),
            .allocator = allocator,
            .virtual_nodes_per_server = virtual_nodes,
        };
    }

    pub fn deinit(self: *ConsistentHash) void {
        for (self.nodes.items) |node| {
            self.allocator.free(node.id);
            self.allocator.free(node.physical_node);
        }
        self.nodes.deinit();
    }

    pub fn addNode(self: *ConsistentHash, id: []const u8) !void {
        var i: u32 = 0;
        while (i < self.virtual_nodes_per_server) : (i += 1) {
            const virtual_id = try std.fmt.allocPrint(self.allocator, "{s} : {d}", .{ id, i });
            const physical_node = try self.allocator.dupe(u8, id);
            const position = try self.hash(virtual_id);
            try self.nodes.append(VirtualNode{
                .id = virtual_id,
                .position = position,
                .physical_node = physical_node,
            });
        }
        std.mem.sort(VirtualNode, self.nodes.items, {}, comptime lessThan);
    }

    pub fn removeNode(self: *ConsistentHash, id: []const u8) void {
        var i: usize = 0;
        while (i < self.nodes.items.len) {
            if (std.mem.eql(u8, self.nodes.items[i].physical_node, id)) {
                const node = self.nodes.orderedRemove(i);
                self.allocator.free(node.id);
                self.allocator.free(node.physical_node);
            } else {
                i += 1;
            }
        }
    }

    pub fn getNode(self: *ConsistentHash, key: []const u8) ![]const u8 {
        if (self.nodes.items.len == 0) {
            return error.NoNodesAvailable;
        }

        const hashValue = try self.hash(key);
        for (self.nodes.items) |node| {
            if (node.position >= hashValue) {
                return node.physical_node;
            }
        }
        return self.nodes.items[0].physical_node;
    }
    
    fn hash(self: *ConsistentHash, key: []const u8) !u64 {
        _ = self;
        var h: u64 = 14695981039346656037;
        for (key) |byte| {
            h = (h ^ byte) *% 1099511628211;
        }
        return h;
    }

    fn lessThan(_: void, a: VirtualNode, b: VirtualNode) bool {
        return a.position < b.position;
    }

    pub fn printRing(self: *ConsistentHash) void {
        std.debug.print("Current Hash Ring State:\n", .{});
        if (self.nodes.items.len == 0) {
            std.debug.print("(empty)\n", .{});
        } else {
            for (self.nodes.items) |node| {
                std.debug.print("Position: {d}, Virtual Node: {s}, Physical Node: {s}\n", .{ node.position, node.id, node.physical_node });
            }
        }
        std.debug.print("\n", .{});
    }

    pub fn visualizeRing(self: *ConsistentHash) !void {
        const radius = 10;
        const center_x = radius + 1;
        const center_y = radius + 1;
        const size = (radius + 1) * 2 + 1;
        
        var circle = try self.allocator.alloc([]u8, size);
        defer self.allocator.free(circle);
        
        for (circle) |*row| {
            row.* = try self.allocator.alloc(u8, size);
            @memset(row.*, ' ');
        }
        defer for (circle) |row| {
            self.allocator.free(row);
        };

        // Draw the circle
        var angle: f64 = 0;
        while (angle < 2 * std.math.pi) : (angle += 0.1) {
            const x = @as(usize, @intFromFloat(@round(center_x + @as(f64, @floatFromInt(radius)) * @cos(angle))));
            const y = @as(usize, @intFromFloat(@round(center_y + @as(f64, @floatFromInt(radius)) * @sin(angle))));
            circle[y][x] = '.';
        }

        // Place nodes on the circle
        for (self.nodes.items) |node| {
            const node_angle = @as(f64, @floatFromInt(node.position)) / @as(f64, std.math.maxInt(u64)) * 2 * std.math.pi;
            const x = @as(usize, @intFromFloat(@round(center_x + @as(f64, @floatFromInt(radius)) * @cos(node_angle))));
            const y = @as(usize, @intFromFloat(@round(center_y + @as(f64, @floatFromInt(radius)) * @sin(node_angle))));
            circle[y][x] = '#';
        }

        // Print the circle
        for (circle) |row| {
            for (row) |cell| {
                std.debug.print("{c}", .{cell});
            }
            std.debug.print("\n", .{});
        }

        // Print node legend
        std.debug.print("\nNodes:\n", .{});
        for (self.nodes.items) |node| {
            std.debug.print("# - {s}\n", .{node.physical_node});
        }
    }
};

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var ch = ConsistentHash.init(allocator, 5); // 5 virtual nodes per server for easier visualization
    defer ch.deinit();

    const stdin = std.io.getStdIn().reader();
    const stdout = std.io.getStdOut().writer();

    const whitespace = " \t\r\n";

    while (true) {
        try stdout.print("\nEnter command (add/remove/get/visualize/quit): ", .{});
        var buf: [100]u8 = undefined;
        if (try stdin.readUntilDelimiterOrEof(buf[0..], '\n')) |user_input| {
            const trimmed = std.mem.trim(u8, user_input, whitespace);
            var it = std.mem.split(u8, trimmed, " ");
            const command = it.next() orelse {
                std.debug.print("No command entered. Current state:\n", .{});
                ch.printRing();
                continue;
            };

            if (std.mem.eql(u8, command, "add")) {
                if (it.next()) |node| {
                    try ch.addNode(node);
                    std.debug.print("Added node: {s}\n", .{node});
                } else {
                    std.debug.print("Please specify a node to add.\n", .{});
                }
            } else if (std.mem.eql(u8, command, "remove")) {
                if (it.next()) |node| {
                    ch.removeNode(node);
                    std.debug.print("Removed node: {s}\n", .{node});
                } else {
                    std.debug.print("Please specify a node to remove.\n", .{});
                }
            } else if (std.mem.eql(u8, command, "get")) {
                if (it.next()) |key| {
                    if (ch.getNode(key)) |node| {
                        std.debug.print("Key '{s}' is handled by node: {s}\n", .{ key, node });
                    } else |_| {
                        std.debug.print("No nodes available.\n", .{});
                    }
                } else {
                    std.debug.print("Please specify a key to get.\n", .{});
                }
            } else if (std.mem.eql(u8, command, "visualize")) {
                try ch.visualizeRing();
            } else if (std.mem.eql(u8, command, "quit")) {
                break;
            } else {
                std.debug.print("Unknown command. Please use add, remove, get, or quit.\n", .{});
            }
            
            std.debug.print("\nCurrent state:\n", .{});
            ch.printRing();
        } else {
            break;
        }
    }
}