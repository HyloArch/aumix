import { useRef } from "react";
import "./Meter.css";
import { useX32Meter } from "../../hooks/useX32";

type MeterProps = {
  typeId: number;
  index: number;
};

export default function Meter({ typeId, index }: MeterProps) {
  const meterRef = useRef<HTMLSpanElement>(null);

  useX32Meter(typeId, index, (level: number) => {
    console.log(level);
    const meterHeight = Math.pow(1.03712367006, 20 * Math.log(level));
    meterRef.current?.style.setProperty(
      "--level-height",
      meterHeight.toString(),
    );
  });

  return <span className="meter" ref={meterRef}></span>;
}
