import { Tabs } from '$sharedComponents/tabs/Tabs';
import { expect, test } from 'vitest';
import { render } from 'vitest-browser-react';

test('Renders nothing if no tabs are provided', () => {
  const { container } = render(
    <Tabs label="Test" activeTabIndex={0} onChangeTab={() => {}} tabs={[]} />,
  );
  expect(container.innerHTML).toBe('');
});
