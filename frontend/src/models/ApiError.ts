export interface IApiError {
  title: string;
  status: number;
  detail: string;
  type: string;
}

export class ApiError extends Error {
  #title: IApiError['title'];
  #status: IApiError['status'];
  #detail: IApiError['detail'];
  #type: IApiError['type'];

  constructor(input: IApiError) {
    super(JSON.stringify(input));
    this.#title = input.title;
    this.#status = input.status;
    this.#detail = input.detail;
    this.#type = input.type;
  }

  get title(): IApiError['title'] {
    return this.#title;
  }

  get status(): IApiError['status'] {
    return this.#status;
  }

  get detail(): IApiError['detail'] {
    return this.#detail;
  }

  get type(): IApiError['type'] {
    return this.#type;
  }

  get entity(): IApiError {
    return {
      title: this.#title,
      status: this.#status,
      detail: this.#detail,
      type: this.#type,
    };
  }

  public isApiError(): this is IApiError {
    return ApiError.isApiError(this);
  }

  public static fromResponse(response: unknown): ApiError | undefined {
    if (!ApiError.isApiError(response)) {
      return undefined;
    }

    return new ApiError(response);
  }

  public static isApiError(input: unknown): input is IApiError {
    if (!input) {
      return false;
    }

    if (!(typeof input === 'object')) {
      return false;
    }

    const hasTitle =
      'title' in input &&
      typeof input.title === 'string' &&
      Boolean(input.title);
    const hasStatus =
      'status' in input &&
      typeof input.status === 'number' &&
      input.status > 99 &&
      input.status < 600;
    const hasDetail =
      'detail' in input &&
      typeof input.detail === 'string' &&
      Boolean(input.detail);
    const hasType =
      'type' in input && typeof input.type === 'string' && Boolean(input.type);
    return hasTitle && hasStatus && hasDetail && hasType;
  }
}
