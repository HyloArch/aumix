import { useRef, useState, type TouchEvent } from "react";
import "./Shows.css";
import type { ShowScene, ShowScene as ShowSceneList } from "../../types";
import { useSocket } from "../../hooks/useSocket";
import {
  EditSceneModal,
  EditShowModal,
  SelectShowModal,
} from "../../components/show-modals";
import { useX32 } from "../../hooks/useX32";
import { clsx } from "../../util";

type ShowModalState = "closed" | "edit" | "new";
type SceneModalState = "closed" | "new" | number;

export default function Shows() {
  const { show } = useX32();

  const [selectShowOpen, setSelectShowOpen] = useState(false);
  const [editShowOpen, setEditShowOpen] = useState<ShowModalState>("closed");
  const [editSceneOpen, setEditSceneOpen] = useState<SceneModalState>("closed");

  return (
    <>
      <div className="shows-page">
        <ShowSidebar
          setSelectShowOpen={setSelectShowOpen}
          setEditShowOpen={setEditShowOpen}
          setEditSceneOpen={setEditSceneOpen}
        />
        <ShowSceneList editScene={setEditSceneOpen} />
      </div>
      <SelectShowModal
        isOpen={selectShowOpen}
        onClose={() => setSelectShowOpen(false)}
      />
      <EditShowModal
        isOpen={editShowOpen != "closed"}
        onClose={() => setEditShowOpen("closed")}
        show={editShowOpen == "edit" ? show : undefined}
      />
      <EditSceneModal
        isOpen={editSceneOpen != "closed"}
        onClose={() => setEditSceneOpen("closed")}
        sceneId={typeof editSceneOpen == "number" ? editSceneOpen : undefined}
      />
    </>
  );
}

type ShowSidebarProps = {
  setSelectShowOpen: (state: boolean) => void;
  setEditShowOpen: (state: ShowModalState) => void;
  setEditSceneOpen: (state: SceneModalState) => void;
};

function ShowSidebar({
  setSelectShowOpen,
  setEditShowOpen,
  setEditSceneOpen,
}: ShowSidebarProps) {
  const { show } = useX32();

  return (
    <>
      <div className="show-sidebar">
        <button onClick={() => setEditShowOpen("new")}>New Show</button>
        <button onClick={() => setSelectShowOpen(true)}>Open Show</button>
        <span className="break"></span>
        <h2>{show?.name}</h2>
        <button disabled={!show} onClick={() => setEditShowOpen("edit")}>
          Edit Show
        </button>
        <button disabled={!show} onClick={() => setEditSceneOpen("new")}>
          Add Scene
        </button>
      </div>
    </>
  );
}

type ShowSceneListProps = {
  editScene: (sceneId: number) => void;
};

type DraggedScene = {
  target: HTMLElement;
  scene: ShowScene;
  index: number;
  top: number;
};

function ShowSceneList({ editScene }: ShowSceneListProps) {
  const { show, currentScene } = useX32();
  const { send } = useSocket();

  const onSceneSelect = (scene: ShowScene) => {
    send({
      op: "SET",
      key: "go-scene",
      value: scene?.id,
    });
  };

  const draggedScene = useRef<DraggedScene>(null);

  const onSceneDrag = (e: TouchEvent, scene: ShowScene, index: number) => {
    if (!draggedScene.current) {
      const target = e.currentTarget as HTMLElement;
      const { top } = target.getBoundingClientRect();
      draggedScene.current = {
        target,
        scene,
        index,
        top: top + 20,
      };
    }
    draggedScene.current.target.style.position = "relative";
    draggedScene.current.target.style.top = `${e.changedTouches[0].pageY - draggedScene.current.top}px`;
  };

  const onSceneRelease = (e: TouchEvent) => {
    if (!draggedScene.current || !show) {
      return;
    }

    let offset = (e.changedTouches[0].pageY - draggedScene.current.top) / 40;
    if (offset < 0) {
      offset += 1;
    }
    const newIndex = Math.max(
      Math.min(draggedScene.current.index + offset, show.scenes.length - 1),
      0,
    );

    send({
      op: "SET",
      key: "scene",
      value: {
        id: draggedScene.current.scene.id,
        newIndex,
      },
    });

    draggedScene.current.target.style.position = "static";
    draggedScene.current = null;
  };

  return (
    <div className="scene-list-container">
      <div className="scene-header">
        <span></span>
        <span>Name</span>
        <span>Movement</span>
        <span>Measure</span>
        <span>X32 Scene</span>
        <span>Sample Count</span>
      </div>
      <div className="scene-list">
        {show?.scenes?.map((scene, index) => (
          <div
            className="scene"
            key={index}
            onClick={() => editScene(index)}
            onTouchMove={(e) => onSceneDrag(e, scene, index)}
            onTouchEnd={onSceneRelease}
          >
            <span
              className={clsx("scene-select", {
                ["selected"]: scene.id == currentScene,
              })}
              onClick={(e) => {
                e.stopPropagation();
                onSceneSelect(scene);
              }}
            />
            <span>{scene.name}</span>
            <span>{scene.movement}</span>
            <span>{scene.measure}</span>
            <span>{scene.sceneId}</span>
            <span>{scene.samples?.length || 0}</span>
          </div>
        ))}
      </div>
    </div>
  );
}
