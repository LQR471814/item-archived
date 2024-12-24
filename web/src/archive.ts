import { createPromiseClient, type PromiseClient } from "@connectrpc/connect"
import { ArchiveService } from "./api/v1/api_connect"
import { createConnectTransport } from "@connectrpc/connect-web"
import type { EntryMetadata } from "./api/v1/api_pb"

const transport = createConnectTransport({
  baseUrl: import.meta.env.VITE_SERVER_URL !== "" ? import.meta.env.VITE_SERVER_URL : window.location.origin
})

export const remote = createPromiseClient(ArchiveService, transport)

export interface Archive {
  read(path: string[]): Promise<{
    metadata: EntryMetadata,
    children?: {
      itemNames: string[]
      containerNames: string[]
    }
  }>
  create(metadata: EntryMetadata, path: string[], createContainer: boolean): Promise<void>
  move(src: string[], dst: string[]): Promise<void>
  delete(path: string[]): Promise<void>
}

export class RemoteArchive implements Archive {
  private client: PromiseClient<typeof ArchiveService>

  constructor(client: PromiseClient<typeof ArchiveService>) {
    this.client = client
  }

  async read(path: string[]) {
    const res = await this.client.read({ path })
    return {
      metadata: res.metadata!,
      children: res.children,
    }
  }
  async create(metadata: EntryMetadata, path: string[], createContainer: boolean): Promise<void> {
    await this.client.create({
      metadata, path, createContainer,
    })
  }
  async move(src: string[], dst: string[]): Promise<void> {
    await this.client.move({
      src, dest: dst,
    })
  }
  async delete(path: string[]): Promise<void> {
    await this.client.delete({ path })
  }

}

export const archive = new RemoteArchive(remote)

