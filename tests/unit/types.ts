export type DataRecord = {
  id: string;
  timestamp: number;
  payload: unknown;
  metadata?: {
    source: string;
    version?: string;
    tags?: string[];
  };
};

export type ParseOptions = {
  strictMode?: boolean;
  validateSchema?: boolean;
  dateFormat?: string;
  maxDepth?: number;
  allowUnknownFields?: boolean;
};

export type ParseResult<T = unknown> = {
  success: boolean;
  data?: T;
  errors?: ParseError[];
  warnings?: string[];
  stats?: {
    parseTimeMs: number;
    recordCount: number;
  };
};

export type ParseError = {
  code: string;
  message: string;
  path?: string;
  value?: unknown;
  severity?: 'error' | 'warning';
};

export type DataSchema = {
  type: 'object' | 'array' | 'string' | 'number' | 'boolean' | 'null';
  properties?: Record<string, DataSchema>;
  items?: DataSchema | DataSchema[];
  required?: string[];
  additionalProperties?: boolean;
  enum?: unknown[];
  format?: string;
  minimum?: number;
  maximum?: number;
  minLength?: number;
  maxLength?: number;
  pattern?: string;
};

export type ParserConfig = {
  name: string;
  description?: string;
  version: string;
  inputFormats: string[];
  outputFormats: string[];
  defaultOptions?: ParseOptions;
  schema?: DataSchema;
};