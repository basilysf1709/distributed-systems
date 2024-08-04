'use client'

import React, { useState, useCallback, useEffect } from 'react';
import ReactFlow, { 
  Controls, 
  Background, 
  useNodesState, 
  useEdgesState,
  addEdge,
  Node,
  Edge,
  Connection
} from 'reactflow';
import 'reactflow/dist/style.css';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Slider } from "@/components/ui/slider";

interface GossipNodeData {
  label: string;
  infected: boolean;
}

const initialNodes: Node<GossipNodeData>[] = [
  { id: '1', position: { x: 100, y: 100 }, data: { label: 'Node 1', infected: false } },
  { id: '2', position: { x: 200, y: 200 }, data: { label: 'Node 2', infected: false } },
  { id: '3', position: { x: 300, y: 100 }, data: { label: 'Node 3', infected: false } },
  { id: '4', position: { x: 400, y: 200 }, data: { label: 'Node 4', infected: false } },
  { id: '5', position: { x: 500, y: 100 }, data: { label: 'Node 5', infected: false } },
];

const initialEdges: Edge[] = [];

const GossipNode: React.FC<{data: GossipNodeData}> = ({ data }) => {
  return (
    <div className={`p-2 rounded-full ${data.infected ? 'bg-red-500' : 'bg-blue-500'}`}>
      {data.label}
    </div>
  );
};

const nodeTypes = {
  gossipNode: GossipNode,
};

interface GossipProtocolVisualizationProps {
  onClose: () => void;
}

export function GossipProtocolVisualization({ onClose }: GossipProtocolVisualizationProps) {
  const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges);
  const [infectionRate, setInfectionRate] = useState<number>(50);
  const [numInfectedNodes, setNumInfectedNodes] = useState<number>(1);
  const [isSimulationRunning, setIsSimulationRunning] = useState<boolean>(false);

  const onConnect = useCallback((params: Connection) => setEdges((eds) => addEdge(params, eds)), [setEdges]);

  const startSimulation = () => {
    setIsSimulationRunning(true);
    setNodes((nds) => nds.map(node => ({ ...node, data: { ...node.data, infected: false } })));
    setNodes((nds) => {
      const newNodes = [...nds];
      for (let i = 0; i < numInfectedNodes; i++) {
        const randomIndex = Math.floor(Math.random() * newNodes.length);
        newNodes[randomIndex] = { ...newNodes[randomIndex], data: { ...newNodes[randomIndex].data, infected: true } };
      }
      return newNodes;
    });
  };

  const stopSimulation = () => {
    setIsSimulationRunning(false);
  };

  useEffect(() => {
    if (isSimulationRunning) {
      const interval = setInterval(() => {
        setNodes((nds) => {
          return nds.map(node => {
            if (node.data.infected) return node;
            const shouldInfect = Math.random() * 100 < infectionRate;
            return shouldInfect ? { ...node, data: { ...node.data, infected: true } } : node;
          });
        });
      }, 1000);
      return () => clearInterval(interval);
    }
  }, [isSimulationRunning, infectionRate]);

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
      <div className="bg-white dark:bg-gray-800 w-11/12 h-5/6 rounded-lg flex">
        <div className="flex-grow">
          <ReactFlow
            nodes={nodes}
            edges={edges}
            onNodesChange={onNodesChange}
            onEdgesChange={onEdgesChange}
            onConnect={onConnect}
            nodeTypes={nodeTypes}
            fitView
          >
            <Controls />
          </ReactFlow>
        </div>
        <Card className="w-80 h-full overflow-y-auto">
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Gossip Protocol Configuration</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <label className="block text-sm font-medium mb-1">Infection Rate (%)</label>
              <Slider 
                value={[infectionRate]} 
                onValueChange={(value) => setInfectionRate(value[0])} 
                max={100} 
                step={1} 
              />
              <span className="text-sm">{infectionRate}%</span>
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Initial Infected Nodes</label>
              <Slider 
                value={[numInfectedNodes]} 
                onValueChange={(value) => setNumInfectedNodes(value[0])} 
                max={nodes.length} 
                step={1} 
              />
              <span className="text-sm">{numInfectedNodes}</span>
            </div>
            <Button onClick={isSimulationRunning ? stopSimulation : startSimulation}>
              {isSimulationRunning ? 'Stop Simulation' : 'Start Simulation'}
            </Button>
            <Button onClick={onClose}>Close</Button>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}