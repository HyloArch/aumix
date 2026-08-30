import { useRef, useState } from "react";
import "./Samples.css";
import { useSocket, useSocketOnMessage } from "../../hooks/useSocket";
import type { Message } from "../../types";

export default function Samples() {
  const { send } = useSocket();

  const [samples, setSamples] = useState<string[]>([]);

  const sampleFileRef = useRef<HTMLInputElement>(null);

  const refreshSamples = () => {
    send({
      op: "GET",
      key: "samples",
    });
  };

  useSocketOnMessage("samples", (message: Message) => {
    if (message.op == "SET") {
      setSamples(message.value as string[]);
    }
  });

  const stopSamples = () => {
    send({
      op: "SET",
      key: "samples",
    });
  };

  const playSample = (sample: string) => {
    send({
      op: "SET",
      key: "samples",
      value: sample,
    });
  };

  const host = window.location.host;

  const deleteSample = async (sample: string) => {
    const response = await fetch(`http://${host}/sample/${sample}`, {
      method: "DELETE",
    });

    if (response.ok) {
      console.log("Sample deleted");
      refreshSamples();
    } else {
      console.error("Server side communication failed:", response.statusText);
    }
  };

  const uploadSample = async () => {
    if (!sampleFileRef.current) {
      return;
    }
    const files = sampleFileRef.current?.files;
    if (!files || files?.length === 0) {
      return;
    }
    const file = files[0];

    const formData = new FormData();
    formData.append("file", file);

    const response = await fetch(`http://${host}/sample/`, {
      method: "POST",
      body: formData,
    });

    if (response.ok) {
      console.log("Upload successfully processed");
      sampleFileRef.current.value = "";
      refreshSamples();
    } else {
      console.error("Server side communication failed:", response.statusText);
    }
  };

  return (
    <div className="samples-container">
      <div className="sample-line">
        <button onClick={refreshSamples}>Refresh Samples</button>
        <button onClick={stopSamples}>Stop Samples</button>
      </div>

      {samples.map((sample) => (
        <div className="sample-line" key={sample}>
          {sample}
          <button onClick={() => playSample(sample)}>Play Sample</button>
          <button onClick={() => deleteSample(sample)}>Delete Sample</button>
        </div>
      ))}

      <div className="sample-line">
        <label htmlFor="sample-file"></label>
        <input
          ref={sampleFileRef}
          name="sample-file"
          type="file"
          accept=".mp3,.wav"
        />
        <button onClick={uploadSample}>Upload Sample</button>
      </div>
    </div>
  );
}
