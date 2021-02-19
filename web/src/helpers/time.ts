export function unixToDate(timestamp: number): string {
  return new Intl.DateTimeFormat("default", {
    year: "numeric",
    month: "numeric",
    day: "numeric",
    hour: "numeric",
    minute: "numeric",
    second: "numeric",
  }).format(new Date(timestamp * 1000));
}

export function durationHumanReadable(totalSeconds: number): string {
  if (totalSeconds === 0) {
    return "< 1s";
  }

  const hours = Math.floor(totalSeconds / 3600);
  const minutes = Math.floor(totalSeconds / 60) % 60;
  const seconds = Math.floor(totalSeconds % 60);

  const hoursString = hours > 0 ? `${hours}h ` : "";
  const minutesString = minutes > 0 ? `${minutes}m ` : "";
  const secondsString = seconds > 0 ? `${seconds}s` : "";

  return `${hoursString}${minutesString}${secondsString}`;
}
