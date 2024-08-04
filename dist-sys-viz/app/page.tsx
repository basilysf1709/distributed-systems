'use client'

import React, { useState, useEffect } from 'react';
import { Button } from "@/components/ui/button";
import { Switch } from "@/components/ui/switch";
import { useTheme } from 'next-themes';
import { Sun, Moon } from 'lucide-react';
import { GossipProtocolVisualization } from '@/components/GossipProtocolVisualization';
import { DistributedSystemsDropdown } from '@/components/DistributedSystemsDropdown';

export default function DistributedSystemsVisualizer() {
  const [showGossipVisualization, setShowGossipVisualization] = useState<boolean>(false);
  const { theme, setTheme } = useTheme();
  const [mounted, setMounted] = useState(false);

  // useEffect to handle mounting
  useEffect(() => {
    setMounted(true);
  }, []);

  // Function to toggle theme
  const toggleTheme = () => {
    setTheme(theme === 'dark' ? 'light' : 'dark');
  };

  // Render nothing until mounted to avoid hydration mismatch
  if (!mounted) return null;

  return (
    <div className="p-4 max-w-4xl mx-auto relative">
      <div className="absolute top-4 right-4 flex items-center space-x-2">
        <Sun className="h-4 w-4" />
        <Switch
          checked={theme === 'dark'}
          onCheckedChange={toggleTheme}
        />
        <Moon className="h-4 w-4" />
      </div>

      <h1 className="text-2xl font-bold mb-8">Distributed Systems Visualizer</h1>
      <DistributedSystemsDropdown />
      <div className="mt-4">
        <Button onClick={() => setShowGossipVisualization(true)}>
          Show Gossip Protocol Visualization
        </Button>
      </div>
      {showGossipVisualization && (
        <GossipProtocolVisualization onClose={() => setShowGossipVisualization(false)} />
      )}
    </div>
  );
}