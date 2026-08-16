import { useRef } from "react";
import { useSocket, useSocketOnMessage } from "../../hooks/useSocket";
import "./Fader.css";
import type { Message } from "../../types";

type FaderProps = {
  type: string;
  id: number;
  onChange?: (level: number) => void;
};

type FaderMessage = {
  type: string;
  index: number;
  level: number;
};

export default function Fader({ type, id, onChange }: FaderProps) {
  const { send } = useSocket();

  const faderRef = useRef<HTMLInputElement>(null);

  const onInput = () => {
    const level = (Number(faderRef.current?.value) || 0) / 1024;
    send({
      op: "SET",
      key: "mix-fader",
      value: {
        type,
        index: id,
        level,
      },
    });
    onChange?.(level);
  };

  const updateFader = (message: Message) => {
    if (faderRef.current === null || !message.value) {
      return;
    }
    const {
      type: mType,
      index: mIndex,
      level: mLevel,
    } = message.value as FaderMessage;
    if (mType === type && mIndex === id) {
      faderRef.current.value = String(mLevel * 1024);
      onChange?.(mLevel);
    }
  };

  useSocketOnMessage("mix-fader", updateFader);

  return (
    <input
      ref={faderRef}
      className="fader"
      type="range"
      min="0"
      max="1024"
      onInput={onInput}
    />
  );
}
