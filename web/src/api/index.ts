import { useContext, useEffect, useState, useCallback } from "react";
import AuthContext from "./AuthContext";
import { useHistory } from "react-router-dom";

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

export const buildApiUrl = (slug: string): string => {
  if (slug.length > 0 && slug[0] === "/") {
    slug = slug.slice(1);
  }

  return `${getAPIURL()}/v1/${slug}`;
};

export const buildWSURL = (slug: string): string => {
  const url = buildApiUrl(slug);

  const parts = url.split("://");
  const protocol = parts[0];

  if (protocol === "http") {
    return `ws://${parts[1]}`;
  }

  return `wss://${parts[1]}`;
};

export const getAPIURL = (): string => {
  return process.env.REACT_APP_API_URL || "http://localhost:8000";
};

const makeRequestInit = (
  authToken: string | null,
  config?: RequestConfig,
  options?: RequestOptions
): RequestInit => {
  const requestInit: RequestInit = {
    method: config?.method ? config.method : HttpMethod.Get,
    headers: config?.headers,
  };

  if (!options?.noAuth && authToken !== null) {
    requestInit.headers = {
      ...requestInit.headers,
      Authorization: `Bearer ${authToken}`,
    };
  }

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

  const auth = useContext(AuthContext);
  const history = useHistory();

  const performRequest = async (): Promise<TData | string> => {
    setError(null);
    setIsLoading(true);

    const response = await fetch(
      buildApiUrl((requestConfig as RequestConfig).slug),
      makeRequestInit(auth.authToken, requestConfig as RequestConfig, options as RequestOptions)
    );

    if (response.status === 401) {
      // TODO: try to refresh token
      history.push("/login");
      auth.setToken(null);
      setIsLoading(false);

      return response.statusText;
    }

    if (!response.ok) {
      setError(response.statusText);
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

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [stringifiedConfig]);

  const execute = useCallback(() => {
    return performRequest();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [stringifiedConfig]);

  return [{ data, isLoading, error }, execute];
};
