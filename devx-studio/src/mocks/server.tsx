import {setupServer} from 'msw/node';
import { cfhandlers } from './continuousfeedback';

export const server = setupServer(...cfhandlers);