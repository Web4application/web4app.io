import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import Home from "../app/page";
import { vi } from "vitest";

// Mock fetch
global.fetch = vi.fn(() =>
  Promise.resolve({ json: () => Promise.resolve({ message: "Assistant Online" }) })
);

describe("Home Page UI", () => {
  it("renders the title and features", () => {
    render(<Home />);
    expect(screen.getByText("Welcome to Bizz")).toBeInTheDocument();
    expect(screen.getByText("Analytics Dashboard")).toBeInTheDocument();
    expect(screen.getByText("Data Upload")).toBeInTheDocument();
  });

  it("launches the AI Assistant on button click", async () => {
    render(<Home />);
    const button = screen.getByText("Launch AI Assistant");
    fireEvent.click(button);

    await waitFor(() => {
      expect(global.fetch).toHaveBeenCalledWith("/api/assistant/start");
    });
  });
});
