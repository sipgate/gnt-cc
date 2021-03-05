export const convertMiBToGiB = (mib: number): number => {
  return Math.round(mib / 102.4) / 10;
};

export const convertMiBToTiB = (mib: number): number => {
  return Math.round(mib / (1024 * 102.4)) / 10;
};

export const prettyPrintMiB = (mib: number): string => {
  if (mib < 1024) {
    return `${mib} MiB`;
  }

  if (mib < 1024 * 1024) {
    return `${convertMiBToGiB(mib)} GiB`;
  }

  return `${convertMiBToTiB(mib)} TiB`;
};
