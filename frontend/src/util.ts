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
