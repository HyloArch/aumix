import { useEffect, useState } from "react";
import "./SelectShowModal.css";
import { useSocket, useSocketOnMessage } from "../../../hooks/useSocket";
import type { Message } from "../../../types";
import Popup from "../../popup";

type SelectShowModalProps = {
  isOpen: boolean;
  onClose: () => void;
};

export default function SelectShowModal({
  isOpen,
  onClose,
}: SelectShowModalProps) {
  const [showList, setShowList] = useState<Map<number, string>>();

  const { send } = useSocket();

  useEffect(() => {
    if (isOpen) {
      send({
        op: "GET",
        key: "show-list",
      });
    }
  }, [isOpen]);

  useSocketOnMessage("show-list", (message: Message) => {
    if (message.op == "SET") {
      const shows = Object.entries(
        message.value as Record<number, string>,
      ).reduce(
        (map, show) => map.set(parseInt(show[0]), show[1]),
        new Map<number, string>(),
      );
      setShowList(shows);
    }
  });

  const selectShow = (id: number) => {
    send({
      op: "SET",
      key: "show",
      value: {
        id,
      },
    });

    onClose();
  };

  const deleteShow = (id: number) => {
    send({
      op: "SET",
      key: "show",
      value: {
        id,
        remove: true,
      },
    });
    send({
      op: "GET",
      key: "show-list",
    });
  };

  return (
    <Popup isOpen={isOpen} onClose={onClose} className="select-show-modal">
      <div className="popup-header">Select Show</div>
      <div className="popup-content">
        {showList &&
          Array.from(showList).map(([id, name]) => (
            <div className="option-row" key={id} onClick={() => selectShow(id)}>
              <span>{name}</span>
              <button
                onClick={(e) => {
                  e.stopPropagation();
                  deleteShow(id);
                }}
              >
                Delete
              </button>
            </div>
          ))}
      </div>
    </Popup>
  );
}
