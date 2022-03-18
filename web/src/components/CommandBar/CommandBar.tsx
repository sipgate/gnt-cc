import React, { ReactElement, useEffect, useRef, useState } from "react";
import { createPortal } from "react-dom";
import { buildApiUrl } from "../../api";
import styles from "./CommandBar.module.scss";

type Instance = {
  clusterName: string;
  name: string;
};

export default function (): ReactElement | null {
  const [isVisible, setIsVisible] = useState(false);
  const [query, setQuery] = useState("");
  const [results, setResults] = useState<Instance[]>([]);

  async function loadResults(query: string) {
    const response = await fetch(buildApiUrl(`search?query=${query}`));

    const data = (await response.json()) as { instances: Instance[] };

    return data.instances;
  }

  useEffect(() => {
    if (query.length > 1) {
      loadResults(query).then((data) => setResults(data));
    } else {
      setResults([]);
    }
  }, [query]);

  useEffect(() => {
    window.addEventListener("click", (e) => {
      setIsVisible(false);
    });
  }, [isVisible]);

  useEffect(() => {
    function handleKeyDown(event: KeyboardEvent) {
      if (
        isModKey(event) &&
        event.key === "k" &&
        event.defaultPrevented === false
      ) {
        event.preventDefault();

        setIsVisible(!isVisible);
      }

      if (event.key === "Escape") {
        if (isVisible) {
          event.stopPropagation();
          setIsVisible(false);
        }
      }
    }

    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [isVisible]);

  if (!isVisible) {
    return null;
  }

  return createPortal(
    <div className={styles.root} onClick={(e) => e.stopPropagation()}>
      <input
        type="text"
        autoFocus
        placeholder="Type something"
        value={query}
        onChange={(e) => setQuery(e.target.value)}
      />
      <div className={styles.results}>
        {results.map((result) => (
          <div key={result.name}>{result.name}</div>
        ))}
      </div>
    </div>,
    document.body
  );
}

export function isModKey(
  event: KeyboardEvent | MouseEvent | React.KeyboardEvent
) {
  return event.ctrlKey;
}
