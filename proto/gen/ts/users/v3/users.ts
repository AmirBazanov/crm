// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               unknown
// source: users/v3/users.proto

/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";

export const protobufPackage = "users.v3";

/** ENUM для стран */
export enum Country {
  COUNTRY_UNSPECIFIED = 0,
  COUNTRY_EN = 1,
  COUNTRY_IT = 2,
  COUNTRY_FR = 3,
  COUNTRY_DE = 4,
  COUNTRY_RU = 5,
  UNRECOGNIZED = -1,
}

/** Сообщение User */
export interface User {
  id: number;
  firstname: string;
  lastname: string;
  nickname: string;
  email: string;
  country: Country;
  createdAt: string;
  updatedAt: string;
}

/** Запрос на создание */
export interface CreateUserRequest {
  firstname: string;
  lastname: string;
  nickname: string;
  email: string;
  password: string;
  country: Country;
}

/** Ответ на создание */
export interface CreateUserResponse {
  id: number;
}

/** Запрос на получение по ID */
export interface GetUserRequest {
  id: number;
}

export interface GetUserResponse {
  user: User | undefined;
}

/** Запрос на обновление */
export interface UpdateUserRequest {
  id: number;
  firstname: string;
  lastname: string;
  nickname: string;
  email: string;
  country: Country;
}

export interface UpdateUserResponse {
  user: User | undefined;
}

/** Запрос на удаление */
export interface DeleteUserRequest {
  id: number;
}

export interface DeleteUserResponse {
  id: number;
}

/** Запрос на получение всех */
export interface GetUsersRequest {
}

export interface GetUsersResponse {
  users: User[];
}

/** Запрос на поиск */
export interface SearchUsersRequest {
  country?: Country | undefined;
  nickname?: string | undefined;
  email?: string | undefined;
  firstname?: string | undefined;
  lastname?: string | undefined;
}

export interface SearchUsersResponse {
  users: User[];
}

export const USERS_V3_PACKAGE_NAME = "users.v3";

/** gRPC-сервис */

export interface UserServiceClient {
  createUser(request: CreateUserRequest): Observable<CreateUserResponse>;

  getUser(request: GetUserRequest): Observable<GetUserResponse>;

  updateUser(request: UpdateUserRequest): Observable<UpdateUserResponse>;

  deleteUser(request: DeleteUserRequest): Observable<DeleteUserResponse>;

  getUsers(request: GetUsersRequest): Observable<GetUsersResponse>;

  search(request: SearchUsersRequest): Observable<SearchUsersResponse>;
}

/** gRPC-сервис */

export interface UserServiceController {
  createUser(
    request: CreateUserRequest,
  ): Promise<CreateUserResponse> | Observable<CreateUserResponse> | CreateUserResponse;

  getUser(request: GetUserRequest): Promise<GetUserResponse> | Observable<GetUserResponse> | GetUserResponse;

  updateUser(
    request: UpdateUserRequest,
  ): Promise<UpdateUserResponse> | Observable<UpdateUserResponse> | UpdateUserResponse;

  deleteUser(
    request: DeleteUserRequest,
  ): Promise<DeleteUserResponse> | Observable<DeleteUserResponse> | DeleteUserResponse;

  getUsers(request: GetUsersRequest): Promise<GetUsersResponse> | Observable<GetUsersResponse> | GetUsersResponse;

  search(
    request: SearchUsersRequest,
  ): Promise<SearchUsersResponse> | Observable<SearchUsersResponse> | SearchUsersResponse;
}

export function UserServiceControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ["createUser", "getUser", "updateUser", "deleteUser", "getUsers", "search"];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod("UserService", method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod("UserService", method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const USER_SERVICE_NAME = "UserService";
