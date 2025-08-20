local M = {}

M.LEVELS = {
	[1] = {
		target = { hash("cup"), hash("shoes"), hash("tutu") },
		pool = {
			{ kind=hash("cup"),    is_correct=true  },
			{ kind=hash("shoes"),  is_correct=true  },
			{ kind=hash("tutu"),   is_correct=true  },
			{ kind=hash("burger"), is_correct=false },
			{ kind=hash("glasses"),is_correct=false },
			{ kind=hash("pizza"),  is_correct=false },
			{ kind=hash("cake"),   is_correct=false },
			{ kind=hash("bean"),   is_correct=false },
		}
	}
}

return M
