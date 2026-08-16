import { MeterMappings } from "../../../constants";
import MixerFader from "../../mixer-fader";
import "./Mixer.css";

export default function Mixer() {
  const meters = [
    MeterMappings.Ch1,
    MeterMappings.Ch2,
    MeterMappings.Ch3,
    MeterMappings.Ch4,
    MeterMappings.Ch5,
    MeterMappings.Ch6,
    MeterMappings.Ch7,
    MeterMappings.Ch8,
  ];

  return (
    <div className="mixer-container">
      {meters.map((m, i) => (
        <MixerFader key={i} mapping={m}></MixerFader>
      ))}
    </div>
  );
}
