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

  return (
    <div className={styles.root} onClick={(e) => e.stopPropagation()}>
      <SearchInput value={query} isLoading={isLoading} onChange={setQuery} />

      <div className={styles.resultsWrapper}>
        {results.clusters.length > 0 && (
          <SearchResults headline="Clusters">
            {results.clusters.map((result) => (
              <SearchResult
                key={result.name}
                name={result.name}
                url={`/${result.name}`}
                onClick={() => setIsVisible(false)}
              />
            ))}
          </SearchResults>
        )}
        {results.instances.length > 0 && (
          <SearchResults headline="Instances">
            {results.instances.map((result) => (
              <SearchResult
                key={result.name}
                name={result.name}
                url={`/${result.clusterName}/instances/${result.name}`}
                onClick={() => setIsVisible(false)}
              />
            ))}
          </SearchResults>
        )}
        {results.nodes.length > 0 && (
          <SearchResults headline="Nodes">
            {results.nodes.map((result) => (
              <SearchResult
                key={result.name}
                name={result.name}
                url={`/${result.clusterName}/nodes/${result.name}`}
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
