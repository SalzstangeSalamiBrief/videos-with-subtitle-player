import { Outlet } from "react-router-dom";
import { Navigation } from "../components/navigation/Navigation";
import { Col, Row } from "antd";
import { FileTreeContextWrapper } from "../contexts/FileTreeContextWrapper";

export function RootLayout() {
  return (
    <FileTreeContextWrapper>
      <Row style={{ height: "100%" }}>
        <Col span={8}>
          <Navigation />
        </Col>
        <Col span={16}>
          <main style={{ padding: "1rem" }}>
            <Outlet />
          </main>
        </Col>
      </Row>
    </FileTreeContextWrapper>
  );
}
