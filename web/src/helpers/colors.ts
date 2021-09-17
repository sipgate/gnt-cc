const COLORS = [
  "#54478c",
  "#2c699a",
  "#048ba8",
  "#0db39e",
  "#16db93",
  "#83e377",
  "#b9e769",
  "#efea5a",
  "#f1c453",
  "#f29e4c",
];

export function getColorForString(groupName: string): string {
  let sum = 0;
  for (let i = 0; i < groupName.length; i++) {
    sum += groupName.charCodeAt(i);
  }

  return COLORS[sum % COLORS.length];
}
