import { useRef, type FocusEvent, type RefObject } from "react";
import type { Sample, ShowScene } from "../../../types";
import "./EditSceneModal.css";
import { useSocket } from "../../../hooks/useSocket";
import Popup from "../../popup";
import { useX32 } from "../../../hooks/useX32";
import { uploadFile } from "../../../util";
import { HOST } from "../../../constants";

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

  const samplesRef = useRef<Sample[]>([]);
  const uploadButtonRef = useRef<HTMLButtonElement>(null);
  const sampleUploadRef = useRef<HTMLInputElement>(null);

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
        samples: samplesRef.current,
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

  const onUpload = () => {
    if (!sampleUploadRef.current || !uploadButtonRef.current) {
      return;
    }

    const files = sampleUploadRef.current?.files;
    if (!files || files?.length !== 1) {
      return;
    }
    const file = files[0];

    uploadButtonRef.current.disabled = true;
    uploadFile(`http://${HOST}/sample/`, file, (percent) => {
      uploadButtonRef.current?.style.setProperty(
        "--upload-percent",
        `${percent}`,
      );
    })
      .then(() => {
        send({
          op: "SET",
          key: "scene",
          value: {
            id: scene?.id,
            samples: [
              ...samplesRef.current,
              {
                name: `Sample ${samplesRef.current.length}`,
                file: file.name,
              },
            ],
          } satisfies Partial<ShowScene>,
        });
      })
      .catch(() => {
        console.log("Error uploading sample");
      });
    uploadButtonRef.current.disabled = false;
  };

  return (
    <Popup isOpen={isOpen} onClose={onClose} className="edit-scene-modal">
      <div className="popup-header">{scene ? "Edit Scene" : "New Scene"}</div>
      <div className="popup-content">
        <div className="scene-info">
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
        </div>
        <SampleList scene={scene} ref={samplesRef} />
        <div className="action-buttons">
          <button onClick={onSave}>Save</button>
          {scene && <button onClick={onDelete}>Delete</button>}
          <button
            className="upload"
            ref={uploadButtonRef}
            onClick={() => sampleUploadRef?.current?.click()}
          >
            Upload Sample
          </button>
          <input
            id="sample-file"
            type="file"
            accept=".mp3,.wav"
            ref={sampleUploadRef}
            hidden
            onChange={onUpload}
          />
        </div>
      </div>
    </Popup>
  );
}

type SampleListProps = {
  scene?: ShowScene;
  ref: RefObject<Sample[]>;
};

function SampleList({ scene, ref }: SampleListProps) {
  const { send } = useSocket();

  ref.current = scene?.samples || [];
  const samples = ref.current;

  const onFocusOut = (e: FocusEvent<HTMLInputElement>, index: number) => {
    ref.current[index].name = e.currentTarget.value;
  };

  const onPlay = (index: number) => {
    index;
  };

  const onDelete = async (index: number) => {
    const response = await fetch(
      `http://${HOST}/sample/${samples[index].file}`,
      {
        method: "DELETE",
      },
    );

    if (response.ok) {
      send({
        op: "SET",
        key: "scene",
        value: {
          id: scene?.id,
          samples: ref.current.filter((_, i) => i !== index),
        } satisfies Partial<ShowScene>,
      });
    } else {
      console.error("Server side communication failed:", response.statusText);
    }
  };

  return (
    <div className="scene-samples">
      <div className="sample-list-header">Samples</div>
      {samples.map((sample, index) => (
        <div className="sample" key={index}>
          <input
            type="text"
            defaultValue={sample.name}
            id="sample-name"
            onBlur={(e: FocusEvent<HTMLInputElement>) => onFocusOut(e, index)}
          />
          <span>{sample.file}</span>
          <button onClick={() => onPlay(index)}>Play</button>
          <button onClick={() => onDelete(index)}>Delete</button>
        </div>
      ))}
    </div>
  );
}
