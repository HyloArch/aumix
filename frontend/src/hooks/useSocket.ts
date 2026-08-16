import { createContext, useContext, useEffect } from "react";
import type { Message } from "../types";

type UseSocketContext = {
  refresh: () => void;
  send: (message: Message) => void;
  close: () => void;
  isConnected: boolean;
  connectHandlers: Set<() => void>;
  messageHandlers: Map<string, Set<(message: Message) => void>>;
  disconnectHandlers: Set<() => void>;
};

export const SocketContext = createContext<UseSocketContext | undefined>(
  undefined,
);

export function useSocket() {
  const context = useContext(SocketContext);

  if (!context) {
    throw new Error("useSocket must be used within a SocketProvider");
  }

  return context;
}

export function useSocketOnConnect(handler: () => void) {
  const { connectHandlers } = useSocket();

  useEffect(() => {
    connectHandlers.add(handler);

    return () => {
      connectHandlers.delete(handler);
    };
  });
}

export function useSocketOnDisconnect(handler: () => void) {
  const { disconnectHandlers } = useSocket();

  useEffect(() => {
    disconnectHandlers.add(handler);

    return () => {
      disconnectHandlers.delete(handler);
    };
  });
}

export function useSocketOnMessage(
  key: string,
  handler: (message: Message) => void,
) {
  const { messageHandlers } = useSocket();

  useEffect(() => {
    let handlers = messageHandlers.get(key);
    if (!handlers) {
      handlers = new Set();
      messageHandlers.set(key, handlers);
    }
    handlers.add(handler);

    return () => {
      handlers.delete(handler);
    };
  }, [handler, key, messageHandlers]);
}
