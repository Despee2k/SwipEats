import { TestBed } from '@angular/core/testing';

import { Match } from './match';

describe('Match', () => {
  let service: Match;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(Match);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
