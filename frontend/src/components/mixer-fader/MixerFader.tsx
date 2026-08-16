import { useRef } from "react";
import type { MeterMapping } from "../../types";
import Fader from "../fader";
import Meter from "../meter";
import "./MixerFader.css";
import { getFaderDB } from "../../x32-utils";

type MixerFaderProps = {
  mapping: MeterMapping;
};

export default function MixerFader({ mapping }: MixerFaderProps) {
  const valueRef = useRef<HTMLDivElement>(null);

  const onChange = (level: number) => {
    if (valueRef.current) {
      valueRef.current.textContent = getFaderDB(level).toString();
    }
  };

  return (
    <div className="mixer-fader">
      <span ref={valueRef} className="fader-value"></span>
      <div className="fader-meter">
        <Fader
          type={mapping.fader.type}
          id={mapping.fader.index}
          onChange={onChange}
        />
        <Meter typeId={mapping.meter.typeId} index={mapping.meter.index} />
      </div>
    </div>
  );
}
