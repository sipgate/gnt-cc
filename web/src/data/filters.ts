const filters: Map<string, Function> = new Map();

filters.set("mbToGb", (value: number): number => {
  return value / 1000;
});

export default filters;
