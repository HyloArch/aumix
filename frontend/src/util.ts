export function clsx(
  ...classes: (string | Record<string, boolean> | undefined)[]
): string {
  return classes
    .flatMap((cls) => {
      if (!cls) {
        return;
      } else if (typeof cls === "string") {
        return cls;
      } else {
        return Object.entries(cls)
          .filter(([_, b]) => b)
          .map(([c]) => c);
      }
    })
    .join(" ");
}

export function uploadFile(
  url: string,
  file: File,
  onProgress: (percent: number) => void,
) {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    const formData = new FormData();

    formData.append("file", file);

    xhr.upload.addEventListener("progress", (event) => {
      if (event.lengthComputable) {
        onProgress(event.loaded / event.total);
      }
    });

    xhr.addEventListener("load", () => {
      if (xhr.status >= 200 && xhr.status < 300) {
        resolve(xhr.responseText);
      } else {
        reject(
          new Error(
            `Upload failed with status: ${xhr.status} ${xhr.statusText}`,
          ),
        );
      }
    });

    xhr.addEventListener("error", () =>
      reject(new Error("Network error occured")),
    );
    xhr.addEventListener("abort", () => reject(new Error("Upload aborted")));

    xhr.open("POST", url, true);
    xhr.send(formData);
  });
}
