import type { IFetchError } from "../types";
import ky from "ky";

export const $api = ky.extend({
  prefixUrl: import.meta.env.VITE_API_URL,
  headers: {
    "Content-Type": "application/json",
  },
  retry: 0,
  timeout: 15000,
  credentials: "include",
  hooks: {
    afterResponse: [
      async (_input, _options, response) => {
        if (!response.ok) {
          const error = await response.json();
          throw FetchError.from(error);
        }
        return response;
      },
    ],
  },
});

export class FetchError extends Error implements IFetchError {
  status_text: string;
  status_code: number;
  message: string;

  static fromUnknown(error: unknown): FetchError {
    if (error instanceof Error) {
      return new FetchError({
        status_text: "Unknown",
        status_code: 500,
        message: error.message,
      });
    }
    if (typeof error === "string") {
      return new FetchError({
        status_text: "Unknown",
        status_code: 500,
        message: error,
      });
    }
    return new FetchError({
      status_text: "Unknown",
      status_code: 500,
      message: "Unknown error occurred.",
    });
  }

  static from(error: unknown) {
    if (typeof error !== "object" || error === null) {
      return FetchError.fromUnknown(error);
    }

    const fetchErr: IFetchError = {
      status_text: "Unknown",
      status_code: 500,
      message: "Unknown error occurred.",
    };

    if ("status_text" in error && typeof error.status_text === "string") {
      fetchErr.status_text = error.status_text;
    }

    if ("status_code" in error && typeof error.status_code === "number") {
      fetchErr.status_code = error.status_code;
    }

    if ("message" in error && typeof error.message === "string") {
      fetchErr.message = error.message;
    }

    return new FetchError(fetchErr);
  }

  constructor(error: IFetchError) {
    super(error.message);
    this.status_text = error.status_text;
    this.status_code = error.status_code;
    this.message = error.message;
  }

  json() {
    return {
      status_text: this.status_text,
      status_code: this.status_code,
      message: this.message,
    };
  }
}
