import {
  useCallback,
  useEffect,
  useRef,
  useState,
  type PropsWithChildren,
} from "react";
import { SocketContext } from "../hooks/useSocket";
import type { Message } from "../types";

const MAX_REFRESHES = 10;

export function SocketProvider({ children }: PropsWithChildren) {
  const socketRef = useRef<WebSocket>(null);
  const refreshTimeout = useRef<number>(null);
  const refreshes = useRef(0);
  const [connectHandlers] = useState<Set<() => void>>(new Set());
  const [messageHandlers] = useState<
    Map<string, Set<(message: Message) => void>>
  >(new Map());
  const [disconnectHandlers] = useState<Set<() => void>>(new Set());

  const [isConnected, setIsSocketConnected] = useState(false);

  const onConnect = useCallback(() => {
    setIsSocketConnected(true);
    connectHandlers.forEach((handler) => handler());
  }, [connectHandlers]);

  const onMessage = useCallback(
    (event: MessageEvent) => {
      const message = JSON.parse(event.data) as Message;

      messageHandlers.get(message.key)?.forEach((handler) => handler(message));
    },
    [messageHandlers],
  );

  const onDisconnect = useCallback(() => {
    setIsSocketConnected(false);
    disconnectHandlers.forEach((handler) => handler());
  }, [disconnectHandlers]);

  const send = (message: Message) => {
    if (socketRef.current?.readyState !== socketRef.current?.OPEN) {
      return;
    }
    const encodedMessage = JSON.stringify(message);
    socketRef.current?.send(encodedMessage);
  };

  const close = () => {
    socketRef.current?.close();
    socketRef.current = null;
  };

  const refresh = useCallback(() => {
    if (
      socketRef.current != null &&
      (socketRef.current?.readyState === socketRef.current?.CONNECTING ||
        socketRef.current?.readyState === socketRef.current?.OPEN)
    ) {
      return;
    }

    const host = window.location.host;
    // const host = "localhost:8080";
    socketRef.current = new WebSocket(`ws://${host}/ws`);
    socketRef.current.onopen = onConnect;
    socketRef.current.onmessage = onMessage;
    socketRef.current.onclose = onDisconnect;
  }, [onConnect, onDisconnect, onMessage]);

  const connectSuccess = () => {
    refreshes.current = 0;

    send({
      op: "GET",
      key: "sync",
    });
  };

  const disconnectRefresh = useCallback(() => {
    if (refreshes.current++ < MAX_REFRESHES) {
      clearTimeout(refreshTimeout.current ?? undefined);
      refreshTimeout.current = setTimeout(() => refresh(), 5000);
    }
  }, [refresh]);

  useEffect(() => {
    if (
      socketRef.current != null &&
      (socketRef.current?.readyState === socketRef.current?.CONNECTING ||
        socketRef.current?.readyState === socketRef.current?.OPEN)
    ) {
      return;
    }

    connectHandlers.add(connectSuccess);
    disconnectHandlers.add(disconnectRefresh);

    refresh();

    return () => {
      close();
      connectHandlers.delete(connectSuccess);
      disconnectHandlers.delete(disconnectRefresh);
    };
  }, [connectHandlers, disconnectHandlers, disconnectRefresh, refresh]);

  return (
    <SocketContext
      value={{
        refresh,
        send,
        close,
        isConnected,
        connectHandlers,
        disconnectHandlers,
        messageHandlers,
      }}
    >
      {children}
    </SocketContext>
  );
}
