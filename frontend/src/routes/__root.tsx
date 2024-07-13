import { Outlet, createRootRoute } from '@tanstack/react-router';
import { Row, Col } from 'antd';
import { Navigation } from '$components/navigation/Navigation';
import { FileTreeContextWrapper } from '$contexts/FileTreeContextWrapper';

export const Route = createRootRoute({
  component: Root,
});

function Root() {
  return (
    <FileTreeContextWrapper>
      <Row style={{ overflow: 'hidden' }}>
        <Col
          span={8}
          style={{
            minHeight: '100lvh',
            maxHeight: '100lvh',
            overflowY: 'auto',
          }}
        >
          <Navigation />
        </Col>
        <Col span={16}>
          <main
            style={{
              maxHeight: '100lvh',
              padding: '1rem',
              overflowY: 'auto',
            }}
          >
            <Outlet />
          </main>
        </Col>
      </Row>
    </FileTreeContextWrapper>
  );
}
