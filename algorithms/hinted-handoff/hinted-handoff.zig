const std = @import("std");

// Struct to represent a data item
const DataItem = struct {
    key: []const u8,
    value: []const u8,
    timestamp: i64,
};

// Struct to represent a node in the system
const Node = struct {
    id: usize,
    data: std.ArrayList(DataItem),
    hints: std.ArrayList(DataItem),
    is_online: bool,

    pub fn init(allocator: std.mem.Allocator, id: usize) Node {
        return Node{
            .id = id,
            .data = std.ArrayList(DataItem).init(allocator),
            .hints = std.ArrayList(DataItem).init(allocator),
            .is_online = true,
        };
    }

    pub fn deinit(self: *Node) void {
        self.data.deinit();
        self.hints.deinit();
    }
};

// Function to simulate writing data to a node
fn writeData(node: *Node, item: DataItem) !void {
    try node.data.append(item);
    std.debug.print("Data written to Node {}: Key: {s}, Value: {s}\n", .{ node.id, item.key, item.value });
}

// Function to simulate data replication
fn replicateData(nodes: []Node, sourceNodeId: usize, item: DataItem) !void {
    for (nodes, 0..) |*node, i| {
        if (i != sourceNodeId) {
            try writeData(node, item);
        }
    }
}

// Function to simulate a node failure and use hinted handoff
fn hintedHandoff(nodes: []Node, targetNodeId: usize, item: DataItem) !void {
    const nextNodeId = (targetNodeId + 1) % nodes.len;
    try nodes[nextNodeId].hints.append(item);
    std.debug.print("Hinted handoff: Data for Node {} stored as hint on Node {}\n", .{ targetNodeId, nextNodeId });
}

// Function to process hints when a node comes back online
fn processHints(node: *Node) !void {
    std.debug.print("Processing hints for Node {}...\n", .{node.id});
    while (node.hints.popOrNull()) |hint| {
        try writeData(node, hint);
        std.debug.print("Processed hint: Key: {s}, Value: {s}\n", .{ hint.key, hint.value });
    }
    std.debug.print("All hints processed for Node {}\n", .{node.id});
}

fn promptEnter() void {
    std.debug.print("\nPress Enter to continue...", .{});
    _ = std.io.getStdIn().reader().readByte() catch {};
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var nodes: [3]Node = undefined;
    for (&nodes, 0..) |*node, i| {
        node.* = Node.init(allocator, i);
    }
    defer for (&nodes) |*node| node.deinit();

    std.debug.print("Hinted Handoff Algorithm Demonstration\n", .{});
    std.debug.print("======================================\n\n", .{});

    std.debug.print("Initial state: Nodes 0, 1, and 2 are online and empty.\n", .{});
    for (nodes, 0..) |node, i| {
        std.debug.print("Node {}: {} data items, {} hints\n", .{ i, node.data.items.len, node.hints.items.len });
    }

    promptEnter();

    // Simulate writing data
    const item = DataItem{
        .key = "example_key",
        .value = "example_value",
        .timestamp = 1628097600,
    };

    std.debug.print("\nStep 1: Writing data to Node 0\n", .{});
    try writeData(&nodes[0], item);

    promptEnter();

    std.debug.print("\nStep 2: Replicating data to Node 1 and Node 2\n", .{});
    try replicateData(&nodes, 0, item);

    std.debug.print("\nCurrent state after replication:\n", .{});
    for (nodes, 0..) |node, i| {
        std.debug.print("Node {}: {} data items, {} hints\n", .{ i, node.data.items.len, node.hints.items.len });
    }

    promptEnter();

    std.debug.print("\nStep 3: Simulating Node 1 going offline\n", .{});
    nodes[1].is_online = false;
    std.debug.print("Node 1 is now offline.\n", .{});

    promptEnter();

    std.debug.print("\nStep 4: Attempting to write new data to Node 1 (which is offline)\n", .{});
    const new_item = DataItem{
        .key = "new_key",
        .value = "new_value",
        .timestamp = 1628097601,
    };
    try hintedHandoff(&nodes, 1, new_item);

    promptEnter();

    std.debug.print("\nStep 5: Checking current state\n", .{});
    for (nodes, 0..) |node, i| {
        std.debug.print("Node {}: {} data items, {} hints, online: {}\n", .{ i, node.data.items.len, node.hints.items.len, node.is_online });
    }

    promptEnter();

    std.debug.print("\nStep 6: Node 1 coming back online\n", .{});
    nodes[1].is_online = true;
    try processHints(&nodes[1]);

    std.debug.print("\nFinal state:\n", .{});
    for (nodes, 0..) |node, i| {
        std.debug.print("Node {}: {} data items, {} hints, online: {}\n", .{ i, node.data.items.len, node.hints.items.len, node.is_online });
    }

    std.debug.print("\nDemonstration complete.\n", .{});
}
