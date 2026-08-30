import { useRef } from "react";
import "./NetworkModal.css";
import { useSocket, useSocketOnMessage } from "../../hooks/useSocket";
import { useX32 } from "../../hooks/useX32";
import { clsx } from "../../util";
import type { Message } from "../../types";
import Popup from "../popup";

type NetworkModalProps = {
  isOpen: boolean;
  close: () => void;
};

export default function NetworkModal({ isOpen, close }: NetworkModalProps) {
  const { isConnected: isSocketConnected, send, refresh } = useSocket();
  const { isConnected: isX32Connected } = useX32();

  useSocketOnMessage("status", (message: Message) => {
    if (message.op === "SET" && message.value) {
      close();
    }
  });

  const ipInputRef = useRef<HTMLInputElement>(null);
  const portInputRef = useRef<HTMLInputElement>(null);

  const connect = () => {
    if (!isSocketConnected) {
      refresh();
      return;
    }

    send({
      op: "SET",
      key: "mixer-address",
      value: {
        ip: ipInputRef.current?.value,
        port: parseInt(portInputRef.current?.value || "0"),
      },
    });
  };

  useSocketOnMessage("mixer-address", (message: Message) => {
    if (message.op === "SET" && ipInputRef.current && portInputRef.current) {
      ipInputRef.current.value = message.value.ip as string;
      portInputRef.current.value = message.value.port as string;
    }
  });

  return (
    <Popup isOpen={isOpen} onClose={close} className="network-modal">
      <div className="popup-header">Network Information</div>
      <div className="popup-content">
        <div className="info-row">
          <span>Web Socket:</span>
          <span className={clsx("status", { connected: isSocketConnected })}>
            {isSocketConnected ? "Connected" : "Disconnected"}
          </span>
        </div>
        <div className="info-row">
          <span>X32 Mixer:</span>
          <span className={clsx("status", { connected: isX32Connected })}>
            {isX32Connected ? "Connected" : "Disconnected"}
          </span>
        </div>
        <div className="info-row">
          <label htmlFor="mixer-ip">Mixer IP: </label>
          <input type="text" id="mixer-ip" ref={ipInputRef} />
        </div>
        <div className="info-row">
          <label htmlFor="mixer-port">Mixer port: </label>
          <input type="number" id="mixer-port" ref={portInputRef} />
        </div>
        <div className="connect-button">
          <button onClick={connect}>Connect</button>
        </div>
      </div>
    </Popup>
  );
}
