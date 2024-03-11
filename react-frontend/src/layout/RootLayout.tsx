import { Outlet } from "react-router-dom";
import { Navigation } from "../components/navigation/Navigation";
import { Col, Row } from "antd";
import { FileTreeContextWrapper } from "../contexts/FileTreeContextWrapper";

export function RootLayout() {
  return (
    <FileTreeContextWrapper>
      <Row style={{ overflow: "hidden" }}>
        <Col
          span={8}
          style={{
            minHeight: "100%",
            maxHeight: "100lvh",
            overflowY: "auto",
          }}
        >
          <Navigation />
        </Col>
        <Col span={16}>
          <main
            style={{
              minHeight: "100%",
              maxHeight: "100lvh",
              padding: "1rem",
              overflowY: "auto",
            }}
          >
            <Outlet />
          </main>
        </Col>
      </Row>
    </FileTreeContextWrapper>
  );
}
