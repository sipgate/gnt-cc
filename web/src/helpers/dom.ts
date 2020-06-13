type ClassNamesObject = { [className: string]: boolean };
type ClassNamesArray = Array<string | ClassNamesObject | null>;

const parseClassNameObject = (inputObject: ClassNamesObject): string => {
  const classes = [];

  for (const name in inputObject) {
    if (inputObject[name] === true) {
      classes.push(name);
    }
  }

  return classes.join(" ");
};

const parseClassNameArray = (inputArray: ClassNamesArray): string => {
  const classes = [];

  for (const element of inputArray) {
    if (typeof element === "string") {
      classes.push(element);
    } else if (element !== null) {
      classes.push(parseClassNameObject(element));
    }
  }

  return classes.join(" ");
};

export const classNameHelper = (
  classNames: string | ClassNamesObject | ClassNamesArray
): string => {
  if (typeof classNames === "string") {
    return classNames;
  }

  if (Array.isArray(classNames)) {
    return parseClassNameArray(classNames);
  }

  return parseClassNameObject(classNames);
};
