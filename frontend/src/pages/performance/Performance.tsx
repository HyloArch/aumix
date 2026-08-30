import { useSocket } from "../../hooks/useSocket";
import { useX32 } from "../../hooks/useX32";
import "./Performance.css";

export default function Performance() {
  const { show, currentScene } = useX32();

  const { send } = useSocket();

  let sceneIndex: number | undefined = undefined;
  const scene = show?.scenes.find((s, i) => {
    if (s.id == currentScene) {
      sceneIndex = i;
      return true;
    }
  });
  const prevSceneIndex =
    show && sceneIndex != undefined && sceneIndex > 0
      ? sceneIndex - 1
      : undefined;
  const nextSceneIndex =
    show &&
    (sceneIndex != undefined
      ? sceneIndex < show?.scenes.length - 1
        ? sceneIndex + 1
        : undefined
      : 0);

  const prevScene = () => {
    if (!show || prevSceneIndex == undefined) {
      return;
    }

    send({
      op: "SET",
      key: "go-scene",
      value: show?.scenes[prevSceneIndex].id,
    });
  };

  const nextScene = () => {
    if (!show || nextSceneIndex == undefined) {
      return;
    }

    send({
      op: "SET",
      key: "go-scene",
      value: show?.scenes[nextSceneIndex].id,
    });
  };

  return (
    <div className="performance-page">
      <div className="show-info">
        {show ? (
          <>
            <div className="show-name">{show?.name}</div>
            <div className="scene-name">{scene?.name}</div>
          </>
        ) : (
          <div className="no-show">No Show Selected</div>
        )}
      </div>
      <div className="scene-actions">
        <button
          className="scene-button"
          onClick={prevScene}
          disabled={prevSceneIndex === undefined}
        >
          <span className="label">Previous</span>
          <span className="value">
            {prevSceneIndex != undefined && show?.scenes[prevSceneIndex].name}
          </span>
        </button>
        <button className="scene-button">
          <span className="label">Play</span>
        </button>
        <button className="scene-button">
          <span className="label">Stop</span>
        </button>
        <button
          className="scene-button"
          onClick={nextScene}
          disabled={nextSceneIndex === undefined}
        >
          <span className="label">Next</span>
          <span className="value">
            {nextSceneIndex != undefined && show?.scenes[nextSceneIndex].name}
          </span>
        </button>
      </div>
    </div>
  );
}
