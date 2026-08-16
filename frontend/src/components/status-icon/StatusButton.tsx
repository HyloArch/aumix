import { useState } from "react";
import { useSocket } from "../../hooks/useSocket";
import { clsx } from "../../util";
import WifiIcon from "../icons/Wifi";
import NetworkModal from "../network-modal";
import "./StatusButton.css";
import { useX32 } from "../../hooks/useX32";

export default function StatusButton() {
  const { isConnected: isSocketConnected, refresh, send } = useSocket();
  const { isConnected: isX32Connected } = useX32();
  const [networkModalOpen, setNetworkModalOpen] = useState(false);

  const onClick = () => {
    refresh();

    send({
      op: "GET",
      key: "mixer-address",
    });

    setNetworkModalOpen(true);
  };

  const close = () => {
    setNetworkModalOpen(false);
  };

  return (
    <>
      <div
        className={clsx("status", "status-button", {
          found: isSocketConnected,
          connected: isX32Connected,
        })}
        onClick={onClick}
      >
        <WifiIcon />
      </div>
      <NetworkModal isOpen={networkModalOpen} close={close} />
    </>
  );
}
