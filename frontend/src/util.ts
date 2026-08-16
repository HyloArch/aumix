export function clsx(...classes: (string | Record<string, boolean>)[]): string {
  return classes.flatMap(cls => {
    if (typeof cls === "string") {
      return cls
    } else {
      return Object.entries(cls).filter(([_, b]) => b).map(([c]) => c)
    }
  }).join(" ")
}