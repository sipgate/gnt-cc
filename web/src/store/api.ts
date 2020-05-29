import PageNames from "@/data/enum/PageNames";
import router, { REDIRECT_TO_QUERY_KEY } from "@/router";
import { Dictionary } from "vue-router/types/router";

export interface LoginCredentials {
  username: string;
  password: string;
}

export type RequestOptions = RequestInit;

export default class Api {
  static readonly tokenStorageKey = "gnt-cc-token";

  static async get(slug: string) {
    const options = {
      method: "GET",
    };

    return Api.request(slug, options);
  }

  static async post(slug: string, body: object) {
    const options = {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json",
      },
    };

    return Api.request(slug, options);
  }

  static async request(slug: string, options: RequestOptions) {
    options = {
      ...options,
      headers: {
        ...options.headers,
        Authorization: `Bearer ${localStorage.getItem(Api.tokenStorageKey)}`,
      },
    };

    const response = await fetch(Api.buildUrl(slug), options);

    if (response.status === 401) {
      const query: Dictionary<string | (string | null)[]> = {};

      if (router.currentRoute.path.replace("/", "") !== "") {
        query[REDIRECT_TO_QUERY_KEY] = router.currentRoute.path;
      }

      await router.push({ name: PageNames.Login, query });
    }

    return response.json();
  }

  static async login(credentials: LoginCredentials) {
    const response = await fetch(Api.buildUrl("login"), {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(credentials),
    });

    return response.json();
  }

  private static buildUrl(slug: string) {
    if (slug.length > 0 && slug[0] === "/") {
      slug = slug.slice(1);
    }

    return `http://localhost:8000/v1/${slug}`;
  }
}
