export function unixToDate(timestamp: number): Date {
  return new Date(timestamp * 1000);
}

export function durationHumanReadable(totalSeconds: number): string {
  if (totalSeconds === 0) {
    return "< 1s";
  }

  const hours = Math.floor(totalSeconds / 3600);
  const minutes = Math.floor(totalSeconds / 60) % 60;
  const seconds = Math.floor(totalSeconds % 60);

  const hoursString = hours > 0 ? `${hours}h` : "";
  const minutesString = minutes > 0 ? `${minutes}m` : "";
  const secondsString = seconds > 0 ? `${seconds}s` : "";

  const hoursSeparator = hours > 0 && (minutes > 0 || seconds > 0) ? " " : "";
  const minutesSeparator = minutes > 0 && seconds > 0 ? " " : "";

  return `${hoursString}${hoursSeparator}${minutesString}${minutesSeparator}${secondsString}`;
}
