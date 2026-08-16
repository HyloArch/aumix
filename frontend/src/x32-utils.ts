export function getFaderDB(level: number) {
  let db = 0;
  if (level == 0) {
    return "-oo";
  } else if (level <= 0.0625) {
    db = -90 + 480 * level;
  } else if (level <= 0.25) {
    db = -60 + 160 * (level - 0.0625);
  } else if (level <= 0.5) {
    db = -30 + 80 * (level - 0.25);
  } else {
    db = -10 + 40 * (level - 0.5);
  }
  return String(db.toFixed(1));
}
