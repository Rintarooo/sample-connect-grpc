import type {
  GetGreetingResponse,
  GetGreetingRequest,
} from "@/generated/ts/grpc/greet/v1/greet_pb";
import { createClient } from "@connectrpc/connect";
import {
  GreetingService,
  GetGreetingRequestSchema,
} from "@/generated/ts/grpc/greet/v1/greet_pb";
import { transport } from "@/api/grpc";
import { create } from "@bufbuild/protobuf";

const greetingClient = createClient(GreetingService, transport);

/**
 * Get a greeting for the specified name using gRPC client
 * @param name - The name to greet
 * @returns Promise<GetGreetingResponse>
 * @throws Error if name is invalid or API call fails
 */
export async function getGreeting(name: string): Promise<GetGreetingResponse> {
  if (!name || name.trim().length === 0) {
    throw new Error("Name is required and cannot be empty");
  }

  try {
    const request: GetGreetingRequest = create(GetGreetingRequestSchema, {
      name: name.trim(),
    });
    const response = await greetingClient.getGreeting(request);

    return response;
  } catch (error) {
    console.error("getGreeting error:", error);
    throw error;
  }
}
