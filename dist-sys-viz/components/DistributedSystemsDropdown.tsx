'use client'

import React, { useState } from 'react';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Button } from "@/components/ui/button";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { ChevronDown } from 'lucide-react';

interface Algorithm {
  name: string;
  subcategories: string[];
}

const algorithms: Algorithm[] = [
  {
    name: "Consensus Algorithms",
    subcategories: ["Paxos", "Raft", "Byzantine Fault Tolerance (BFT)", "Zab (used in Apache ZooKeeper)"]
  },
  {
    name: "Leader Election",
    subcategories: ["Bully Algorithm", "Ring Algorithm", "Token Ring"]
  },
  {
    name: "Clock Synchronization",
    subcategories: ["Network Time Protocol (NTP)", "Berkeley Algorithm", "Cristian's Algorithm", "Vector Clocks", "Lamport Timestamps"]
  },
  {
    name: "Replication and Consistency",
    subcategories: ["Primary-Backup Replication", "Chain Replication", "Quorum-based Replication", "Eventual Consistency", "Strong Consistency"]
  },
  {
    name: "Distributed Transactions",
    subcategories: ["Two-Phase Commit (2PC)", "Three-Phase Commit (3PC)", "Saga Pattern"]
  },
  {
    name: "Partitioning and Sharding",
    subcategories: ["Consistent Hashing", "Range Partitioning", "Hash Partitioning"]
  },
  {
    name: "Load Balancing",
    subcategories: ["Round Robin", "Least Connections", "IP Hash"]
  },
  {
    name: "Failure Detection and Recovery",
    subcategories: ["Heartbeat Mechanism", "Gossip Protocol", "Phi Accrual Failure Detector"]
  },
  {
    name: "Distributed File Systems",
    subcategories: ["Google File System (GFS) architecture", "Hadoop Distributed File System (HDFS)"]
  },
  {
    name: "Distributed Databases",
    subcategories: ["CAP Theorem visualization", "PACELC Theorem", "Read Repair in NoSQL databases"]
  },
  {
    name: "Distributed Caching",
    subcategories: ["Cache Coherence Protocols", "Distributed Hash Tables (DHT)"]
  },
  {
    name: "Message Queuing and Streaming",
    subcategories: ["Publish-Subscribe Model", "Kafka Architecture"]
  },
  {
    name: "Service Discovery",
    subcategories: ["Client-Side Discovery", "Server-Side Discovery"]
  },
  {
    name: "Distributed Tracing",
    subcategories: ["Zipkin Architecture", "Jaeger Architecture"]
  },
  {
    name: "Distributed Locking",
    subcategories: ["Redlock Algorithm (Redis)", "Chubby (Google)"]
  },
  {
    name: "Peer-to-Peer Networks",
    subcategories: ["BitTorrent Protocol", "Chord Protocol"]
  },
  {
    name: "Distributed Graph Processing",
    subcategories: ["Pregel Model (used in Google's PageRank)"]
  },
  {
    name: "Eventual Consistency Techniques",
    subcategories: ["Conflict-free Replicated Data Types (CRDTs)", "Operational Transformation"]
  }
];

export function DistributedSystemsDropdown() {
  const [selectedCategory, setSelectedCategory] = useState<string | undefined>();
  const [selectedSubcategory, setSelectedSubcategory] = useState<string | undefined>();

  return (
    <div className="w-full space-y-2">
      <Select onValueChange={setSelectedCategory}>
        <SelectTrigger className="w-full">
          <SelectValue placeholder="Select a category" />
        </SelectTrigger>
        <SelectContent>
          {algorithms.map((algo) => (
            <SelectItem key={algo.name} value={algo.name}>
              {algo.name}
            </SelectItem>
          ))}
        </SelectContent>
      </Select>

      {selectedCategory && (
        <Popover>
          <PopoverTrigger asChild>
            <Button variant="outline" className="w-full justify-between">
              {selectedSubcategory || "Select a subcategory"}
              <ChevronDown className="h-4 w-4 opacity-50" />
            </Button>
          </PopoverTrigger>
          <PopoverContent className="w-full p-0">
            <div className="max-h-[300px] overflow-y-auto">
              {algorithms.find(algo => algo.name === selectedCategory)?.subcategories.map((sub) => (
                <Button
                  key={sub}
                  variant="ghost"
                  className="w-full justify-start font-normal"
                  onClick={() => setSelectedSubcategory(sub)}
                >
                  {sub}
                </Button>
              ))}
            </div>
          </PopoverContent>
        </Popover>
      )}

      {selectedSubcategory && (
        <div className="p-4 border rounded-md">
          Selected: {selectedCategory} - {selectedSubcategory}
        </div>
      )}
    </div>
  );
}