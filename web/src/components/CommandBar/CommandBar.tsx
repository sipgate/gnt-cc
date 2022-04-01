import React, { ReactElement, useEffect, useState } from "react";
import { buildApiUrl } from "../../api";
import styles from "./CommandBar.module.scss";
import SearchInput from "./SearchInput/SearchInput";
import SearchResult from "./SearchResult/SearchResult";
import SearchResults from "./SearchResults/SearchResults";

type ClusterSearchResult = {
  name: string;
};

type ResourceSearchResult = {
  clusterName: string;
  name: string;
};

type SearchResultsResponse = {
  clusters: ClusterSearchResult[];
  instances: ResourceSearchResult[];
  nodes: ResourceSearchResult[];
};

const initialState = {
  clusters: [],
  instances: [],
  nodes: [],
};

function useSearchResults(): [
  {
    query: string;
    results: SearchResultsResponse;
    isLoading: boolean;
  },
  (query: string) => void
] {
  const [query, setQuery] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [results, setResults] = useState<SearchResultsResponse>(initialState);

  async function loadResults(
    query: string,
    abortController: AbortController
  ): Promise<SearchResultsResponse> {
    const response = await fetch(buildApiUrl(`search?query=${query}`), {
      signal: abortController.signal,
    });

    return response.json();
  }

  useEffect(() => {
    const abortController = new AbortController();

    if (query.length > 0) {
      setIsLoading(true);
      loadResults(query, abortController)
        .then((data) => {
          setResults(data);
          setIsLoading(false);
        })
        .catch(() => {
          setIsLoading(false);
        });
    } else {
      setResults(initialState);
    }

    return () => {
      abortController.abort();
    };
  }, [query]);

  return [
    {
      query,
      results,
      isLoading,
    },
    setQuery,
  ];
}

export default function (): ReactElement | null {
  const [isVisible, setIsVisible] = useState(false);
  const [selectionIndex, setSelectionIndex] = useState(0);

  const [{ query, results, isLoading }, setQuery] = useSearchResults();

  useEffect(() => {
    window.addEventListener("click", () => {
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

      if (!isVisible) {
        return;
      }

      if (event.key === "Escape") {
        event.stopPropagation();
        setIsVisible(false);
      }

      const { clusters, instances, nodes } = results;
      const totalResults = clusters.length + instances.length + nodes.length;

      if (totalResults === 0) {
        return;
      }

      if (event.key === "ArrowDown") {
        event.preventDefault();
        setSelectionIndex((selectionIndex) => {
          const index = selectionIndex + 1;
          return index === totalResults ? 0 : index;
        });
      }

      if (event.key === "ArrowUp") {
        event.preventDefault();
        setSelectionIndex((selectionIndex) => {
          const index = selectionIndex - 1;
          return index === -1 ? totalResults - 1 : index;
        });
      }

      if (event.key === "Enter") {
        document.querySelector<HTMLElement>('[data-selected="true"]')?.click();
      }
    }

    setSelectionIndex(0);

    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [isVisible, results]);

  if (!isVisible) {
    return null;
  }

  const { clusters, instances, nodes } = results;

  return (
    <div className={styles.root} onClick={(e) => e.stopPropagation()}>
      <SearchInput value={query} isLoading={isLoading} onChange={setQuery} />

      <div className={styles.resultsWrapper}>
        {instances.length > 0 && (
          <SearchResults headline="Instances">
            {instances.map((result, i) => (
              <SearchResult
                key={result.name}
                name={result.name}
                url={`/${result.clusterName}/instances/${result.name}`}
                selected={i === selectionIndex}
                onClick={() => setIsVisible(false)}
              />
            ))}
          </SearchResults>
        )}
        {clusters.length > 0 && (
          <SearchResults headline="Clusters">
            {clusters.map((result, i) => (
              <SearchResult
                key={result.name}
                name={result.name}
                url={`/${result.name}`}
                selected={i + instances.length === selectionIndex}
                onClick={() => setIsVisible(false)}
              />
            ))}
          </SearchResults>
        )}
        {nodes.length > 0 && (
          <SearchResults headline="Nodes">
            {nodes.map((result, i) => (
              <SearchResult
                key={result.name}
                name={result.name}
                url={`/${result.clusterName}/nodes/${result.name}`}
                selected={
                  i + instances.length + clusters.length === selectionIndex
                }
                onClick={() => setIsVisible(false)}
              />
            ))}
          </SearchResults>
        )}
      </div>
    </div>
  );
}

export function isModKey(
  event: KeyboardEvent | MouseEvent | React.KeyboardEvent
) {
  return event.ctrlKey;
}
