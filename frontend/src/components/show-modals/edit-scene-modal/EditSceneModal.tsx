import { useRef } from "react";
import type { ShowScene } from "../../../types";
import "./EditSceneModal.css";
import { useSocket } from "../../../hooks/useSocket";
import Popup from "../../popup";
import { useX32 } from "../../../hooks/useX32";

type EditSceneModalProps = {
  isOpen: boolean;
  onClose: () => void;
  sceneId?: number;
};

export default function EditSceneModal({
  sceneId,
  isOpen,
  onClose,
}: EditSceneModalProps) {
  const nameInputRef = useRef<HTMLInputElement>(null);
  const movementInputRef = useRef<HTMLInputElement>(null);
  const measureInputRef = useRef<HTMLInputElement>(null);
  const x32SceneInputRef = useRef<HTMLInputElement>(null);

  const { show } = useX32();

  const scene = sceneId != undefined ? show?.scenes[sceneId] : undefined;

  const { send } = useSocket();

  const onSave = () => {
    if (
      !nameInputRef.current ||
      !movementInputRef.current ||
      !measureInputRef.current ||
      !x32SceneInputRef.current
    ) {
      return;
    }

    send({
      op: "SET",
      key: "scene",
      value: {
        id: scene?.id,
        name: nameInputRef.current.value,
        movement: movementInputRef.current.valueAsNumber,
        measure: measureInputRef.current.valueAsNumber,
        sceneId: x32SceneInputRef.current.valueAsNumber,
      } satisfies Partial<ShowScene>,
    });

    onClose();
  };

  const onDelete = () => {
    send({
      op: "SET",
      key: "scene",
      value: {
        id: scene?.id,
        remove: true,
      },
    });

    onClose();
  };

  return (
    <Popup isOpen={isOpen} onClose={onClose} className="edit-scene-modal">
      <div className="popup-header">{scene ? "Edit Scene" : "New Scene"}</div>
      <div className="popup-content">
        <div className="info-row">
          <label htmlFor="scene-name">Name:</label>
          <input
            type="text"
            id="scene-name"
            ref={nameInputRef}
            defaultValue={scene?.name}
          />
        </div>
        <div className="info-row">
          <label htmlFor="scene-movement">Movement:</label>
          <input
            type="number"
            id="scene-movement"
            ref={movementInputRef}
            defaultValue={scene?.movement}
          />
        </div>
        <div className="info-row">
          <label htmlFor="scene-measure">Measure:</label>
          <input
            type="number"
            id="scene-measure"
            ref={measureInputRef}
            defaultValue={scene?.measure}
          />
        </div>
        <div className="info-row">
          <label htmlFor="x32-scene">X32 Scene Id:</label>
          <input
            type="number"
            id="x32-scene"
            ref={x32SceneInputRef}
            defaultValue={scene?.sceneId}
          />
        </div>
        <div className="action-buttons">
          <button onClick={onSave}>Save</button>
          {scene && <button onClick={onDelete}>Delete</button>}
        </div>
      </div>
    </Popup>
  );
}
