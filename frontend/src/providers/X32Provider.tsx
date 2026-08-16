import { useState, type PropsWithChildren } from "react";
import { useSocketOnDisconnect, useSocketOnMessage } from "../hooks/useSocket";
import { getHandlerKey, X32Context } from "../hooks/useX32";
import type { Message } from "../types";

export function X32SProvider({ children }: PropsWithChildren) {
  const [isConnected, setIsConnected] = useState(false);
  const [meterHandlers] = useState<Map<string, Set<(level: number) => void>>>(
    new Map(),
  );

  useSocketOnMessage("status", (message: Message) => {
    if (message.op === "SET") {
      setIsConnected(Boolean(message.value));
    }
  });

  useSocketOnDisconnect(() => {
    setIsConnected(false);
  });

  useSocketOnMessage("meters", (message: Message) => {
    if (message.op !== "SET") {
      return;
    }

    const meterIndex = message.value.index as number;
    const length = message.value.length as number;
    const levelBytes = Uint8Array.from(atob(message.value.levels), (ch) =>
      ch.charCodeAt(0),
    );
    const view = new DataView(levelBytes.buffer);
    const levels = Array.from(new Array(length), (_, i) =>
      view.getFloat32(i, true),
    );

    levels.forEach((level, index) => {
      const handlers = meterHandlers.get(getHandlerKey(meterIndex, index));
      handlers?.forEach((handler) => handler(level));
    });
  });

  return (
    <X32Context value={{ isConnected, meterHandlers }}>{children}</X32Context>
  );
}
