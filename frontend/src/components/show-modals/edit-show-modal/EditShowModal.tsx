import { useRef } from "react";
import type { Show } from "../../../types";
import "./EditShowModal.css";
import { useSocket } from "../../../hooks/useSocket";
import Popup from "../../popup";

type EditShowModalProps = {
  isOpen: boolean;
  onClose: () => void;
  show?: Show;
};

export default function EditShowModal({
  isOpen,
  onClose,
  show,
}: EditShowModalProps) {
  const nameInputRef = useRef<HTMLInputElement>(null);

  const { send } = useSocket();

  const onSave = () => {
    if (!nameInputRef.current || nameInputRef.current?.value == "") {
      return;
    }

    send({
      op: "SET",
      key: "show",
      value: {
        id: show?.id,
        name: nameInputRef.current.value,
      } satisfies Partial<Show>,
    });
    onClose();
  };

  return (
    <Popup isOpen={isOpen} onClose={onClose} className="edit-show-modal">
      <div className="popup-header">{show ? "Edit Show" : "New Show"}</div>
      <div className="popup-content">
        <div className="info-row">
          <label htmlFor="show-name">Name:</label>
          <input
            type="text"
            id="show-name"
            ref={nameInputRef}
            defaultValue={show?.name}
          />
        </div>
        <div className="action-buttons">
          <button onClick={onSave}>Save</button>
        </div>
      </div>
    </Popup>
  );
}
