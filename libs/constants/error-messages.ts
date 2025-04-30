// src/common/constants/error-messages.ts

export const PRISMA_ERROR_MESSAGES = {
  UNIQUE_CONSTRAINT_VIOLATION: 'Unique constraint violation',
  FOREIGN_KEY_CONSTRAINT_VIOLATION: 'Foreign key constraint violation',
  RECORD_NOT_FOUND: 'Record not found',
  INTERNAL_SERVER_ERROR: 'Internal server error while processing the request.',
  INVALID_DATA_PROVIDED: 'Invalid data provided, please check the request.',
  UNEXPECTED_ERROR: 'Unexpected error occurred: ',
  UNIQUE_CONSTRAINT_VIOLATION_DETAILS: (target: unknown) =>
    `Unique constraint violation: ${target}`,
  FOREIGN_KEY_CONSTRAINT_VIOLATION_DETAILS: (target: unknown) =>
    `Foreign key constraint violation: ${target}`,
  RECORD_NOT_FOUND_DETAILS: (model: unknown) => `Record not found for ${model}`,
  CONNECTION_ERROR: 'Database connection error',
};

export const PRISMA_MESSAGES = {
  connectionSuccess: 'Prisma connected to the database successfully',
  connectionError: 'Error connecting to the database',
};
