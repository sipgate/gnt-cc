import React from "react";
import {
  findByText as findByTextGlobal,
  fireEvent,
  Matcher,
  render,
  waitFor,
} from "@testing-library/react";
import each from "jest-each";
import { rest } from "msw";
import { setupServer } from "msw/node";
import InstanceActions from "./InstanceActions";
import { GntInstance } from "../../api/models";
import JobWatchContext from "../../contexts/JobWatchContext";

type Actions = "failover" | "migrate" | "start" | "restart" | "shutdown";

const jobIds = {
  failover: 421,
  migrate: 422,
  start: 423,
  restart: 424,
  shutdown: 425,
};

const server = setupServer(
  rest.post<null, { action: Actions }>(
    "/api/v1/clusters/testClusterName/instances/testInstance/:action",
    (req, res, ctx) => {
      const jobId = jobIds[req.params.action];

      if (!jobId) {
        return res(ctx.status(400), ctx.body("invalid action"));
      }

      return res(
        ctx.json({
          jobId,
        })
      );
    }
  )
);

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

function createMockInstance(overrideParams: Partial<GntInstance>): GntInstance {
  return {
    name: "testInstance",
    cpuCount: 1,
    disks: [],
    isRunning: true,
    memoryTotal: 1024,
    nics: [],
    offersVnc: false,
    primaryNode: "",
    secondaryNodes: [],
    tags: [],
    ...overrideParams,
  };
}

each([
  ["failover", /^failover$/i, createMockInstance({ isRunning: true })],
  ["migrate", /^migrate$/i, createMockInstance({ isRunning: true })],
  ["start", /^start$/i, createMockInstance({ isRunning: false })],
  ["restart", /^restart$/i, createMockInstance({ isRunning: true })],
  ["shutdown", /^shutdown$/i, createMockInstance({ isRunning: true })],
]).test(
  "%s action button triggers corresponding action",
  async (name: Actions, matcher: Matcher, instance: GntInstance) => {
    const mockTrackJob = jest.fn();

    const { findByText, findByRole } = render(
      <JobWatchContext.Provider
        value={{
          trackedJobs: [],
          trackJob: mockTrackJob,
          untrackJob: jest.fn(),
        }}
      >
        <InstanceActions clusterName="testClusterName" instance={instance} />
      </JobWatchContext.Provider>
    );

    const button = await findByText(matcher);

    fireEvent.click(button);

    const dialog = await findByRole("dialog");
    const confirmButton = await findByTextGlobal(dialog, matcher);

    fireEvent.click(confirmButton);

    await waitFor(() => {
      expect(mockTrackJob).toHaveBeenCalledWith(jobIds[name]);
    });
  }
);
