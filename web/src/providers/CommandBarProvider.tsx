import React, { ReactElement, PropsWithChildren, useState } from "react";
import SearchBarContext from "../contexts/SearchBarContext";

export default function CommandBarProvider({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  const [isVisible, setIsVisible] = useState(false);

  function toggleVisibility() {
    setIsVisible(!isVisible);
  }

  function setVisible(visible: boolean) {
    setIsVisible(visible);
  }

  return (
    <SearchBarContext.Provider
      value={{
        isVisible,
        toggleVisibility,
        setVisible,
      }}
    >
      {children}
    </SearchBarContext.Provider>
  );
}
