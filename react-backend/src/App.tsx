import { BrowserRouter, Route, Routes } from "react-router-dom";
import { RootLayout } from "./layout/RootLayout";
import { AudioFilePage } from "./pages/AudioFile/AudioFilePage";

export function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<RootLayout />}>
          <Route path="audio-file/:audioFileId" element={<AudioFilePage />} />
        </Route>
        {/* TODO NOT FOUND */}
      </Routes>
    </BrowserRouter>
  );
}
