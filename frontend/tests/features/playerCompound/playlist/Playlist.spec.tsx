/**
 * ! Currently components that use tanstack router cannot be tested
 * ! an un fixable error is thrown when the content of the file is commented in:
 * !  TypeError: Cannot read properties of undefined (reading 'update')
 * !    src/routeTree.gen.ts:49:53
 */

// import { Playlist } from '$features/playerCompound/playlist/Playlist';
// import { screen, within } from '@testing-library/react';
// import { describe, expect, test } from 'vitest';
// import { RenderFakeRouterShell } from '../../../utlities/RenderFakeRouterShell';

// // COMPONENTS THAT USE TANSTACK ROUTER CANNOT BE TESTED AT THE CURRENT STATE
// describe.skip('Playlist', () => {
//   test('Render only header if no siblings exist', () => {
//     RenderFakeRouterShell(() => <Playlist siblings={[]} />);
//     expect(screen.getByText('Playlist')).toBeInTheDocument();
//     const list = screen.queryByRole('list');
//     expect(list).toBeInTheDocument();
//     const listItems = within(list!).queryAllByRole('listitem');
//     expect(listItems.length).toBe(0);
//   });
// });
