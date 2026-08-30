import { createContext, useContext, useEffect } from "react";
import type { Show } from "../types";

type UseX32Context = {
  isConnected: boolean;
  meterHandlers: Map<string, Set<(level: number) => void>>;
  show?: Show;
  currentScene?: number;
};

export const X32Context = createContext<UseX32Context | undefined>(undefined);

export function useX32() {
  const context = useContext(X32Context);

  if (!context) {
    throw new Error("useX32 must be used within a X32Provider");
  }

  return context;
}

export function useX32Meter(
  typeId: number,
  index: number,
  handler: (level: number) => void,
) {
  const { meterHandlers } = useX32();

  useEffect(() => {
    const key = getHandlerKey(typeId, index);
    let handlers = meterHandlers.get(key);
    if (!handlers) {
      handlers = new Set();
      meterHandlers.set(key, handlers);
    }
    handlers.add(handler);

    return () => {
      handlers.delete(handler);
    };
  }, [meterHandlers]);
}

export function getHandlerKey(typeId: number, index: number) {
  return `${typeId}__${index}`;
}
