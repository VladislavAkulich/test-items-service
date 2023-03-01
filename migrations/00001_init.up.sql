CREATE TABLE IF NOT EXISTS public.test_case
(
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name          TEXT NOT NULL,
    steps         TEXT[],
    preconditions TEXT,
    author        TEXT NOT NULL
);

-- DATA --

INSERT INTO public.test_case (name, steps, preconditions, author)
VALUES ('test case 1', ARRAY ['step1', 'step2'], 'preconditions 1', 'gleb');
INSERT INTO public.test_case (name, steps, preconditions, author)
VALUES ('test case 2', ARRAY ['step3', 'step4'], 'preconditions 2', 'vladik');
