import { useCallback, useEffect, useState } from "react";

export enum HttpMethod {
  Get = "GET",
  Post = "POST",
  Put = "PUT",
  Delete = "DELETE",
  Patch = "PATCH",
}

export interface RequestConfig {
  headers?: HeadersInit;

  body?: BodyInit | null;

  slug: string;

  method?: HttpMethod;
}

const stripLeadingSlug = (slug: string): string => {
  if (slug.length > 0 && slug[0] === "/") {
    return slug.slice(1);
  }

  return slug;
};

export const buildApiUrl = (slug: string): string => {
  return `/api/v1/${stripLeadingSlug(slug)}`;
};

export const buildWSURL = (slug: string): string => {
  const { hostname, port, protocol } = window.location;

  const wsProtocol = protocol === "http:" ? "ws:" : "wss:";
  return `${wsProtocol}//${hostname}:${port}/api/v1/${stripLeadingSlug(slug)}`;
};

const makeRequestInit = (config?: RequestConfig): RequestInit => {
  const requestInit: RequestInit = {
    method: config?.method ? config.method : HttpMethod.Get,
    headers: config?.headers,
  };

  if (config?.body) {
    requestInit.body = config.body;
    requestInit.headers = {
      ...requestInit.headers,
      "Content-type": "application/json",
    };
  }

  return requestInit;
};

interface UseApiState<TData> {
  data: TData | null;
  isLoading: boolean;
  error: string | null;
}

interface RequestOptions {
  noAuth?: boolean;
  manual?: boolean;
}

export const useApi = <TData>(
  requestConfig: RequestConfig | string,
  options?: RequestOptions
): [UseApiState<TData>, () => Promise<TData | string>] => {
  if (typeof requestConfig === "string") {
    requestConfig = {
      slug: requestConfig,
    };
  }

  options = {
    manual: false,
    noAuth: false,
    ...options,
  };

  const stringifiedConfig = JSON.stringify(requestConfig);

  const [data, setData] = useState<TData | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const performRequest = async (): Promise<TData | string> => {
    setError(null);
    setIsLoading(true);

    const response = await fetch(
      buildApiUrl((requestConfig as RequestConfig).slug),
      makeRequestInit(requestConfig as RequestConfig)
    );

    if (response.status === 401) {
      // TODO: try to refresh token
      setIsLoading(false);

      return response.statusText;
    }

    if (!response.ok) {
      const body = await response.json();

      if (isErrorBody(body)) {
        setError(body.error);
      } else {
        setError(response.statusText);
      }

      setIsLoading(false);

      return response.statusText;
    }

    const data = (await response.json()) as TData;
    setData(data);
    setIsLoading(false);
    return data;
  };

  useEffect(() => {
    if (!options?.manual) {
      performRequest();
    }
  }, [stringifiedConfig]);

  const execute = useCallback(() => {
    return performRequest();
  }, [stringifiedConfig]);

  return [{ data, isLoading, error }, execute];
};

function isErrorBody(body: unknown): body is { error: string } {
  return (
    body !== null &&
    typeof body === "object" &&
    typeof (body as { error: string }).error === "string"
  );
}
