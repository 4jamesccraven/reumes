package resume

import (
	"github.com/noirbizarre/gonja"
)

func toContext[T any](in []T, f func(T) map[string]any) []map[string]any {
	out := make([]map[string]any, len(in))

	for i, item := range in {
		out[i] = f(item)
	}

	return out
}

func (self *Resume) baseContext() gonja.Context {
	return gonja.Context{
		"basics": map[string]any{
			"name":    self.Basics.Name,
			"label":   self.Basics.Label,
			"image":   self.Basics.Image,
			"email":   self.Basics.Email,
			"phone":   self.Basics.Phone,
			"url":     self.Basics.Url,
			"summary": self.Basics.Summary,
			"location": map[string]any{
				"address":     self.Basics.Location.Address,
				"postalCode":  self.Basics.Location.PostalCode,
				"city":        self.Basics.Location.City,
				"countryCode": self.Basics.Location.CountryCode,
				"region":      self.Basics.Location.Region,
			},
			"profiles": toContext(self.Basics.Profiles, func(p Profile) map[string]any {
				return map[string]any{
					"network":  p.Network,
					"username": p.Username,
					"url":      p.Url,
				}
			}),
		},
		"work": toContext(self.Work, func(w Work) map[string]any {
			return map[string]any{
				"name":       w.Name,
				"position":   w.Position,
				"url":        w.Url,
				"startDate":  w.StartDate,
				"endDate":    w.EndDate,
				"summary":    w.Summary,
				"highlights": w.Highlights,
			}
		}),
		"volunteer": toContext(self.Volunteer, func(v Volunteer) map[string]any {
			return map[string]any{
				"organization": v.Organization,
				"position":     v.Position,
				"url":          v.Url,
				"startDate":    v.StartDate,
				"endDate":      v.EndDate,
				"summary":      v.Summary,
				"highlights":   v.Highlights,
			}
		}),
		"education": toContext(self.Education, func(e Education) map[string]any {
			return map[string]any{
				"institution": e.Institution,
				"url":         e.Url,
				"area":        e.Area,
				"studyType":   e.StudyType,
				"startDate":   e.StartDate,
				"endDate":     e.EndDate,
				"score":       e.Score,
				"courses":     e.Courses,
			}
		}),
		"awards": toContext(self.Awards, func(a Award) map[string]any {
			return map[string]any{
				"title":   a.Title,
				"awarder": a.Awarder,
				"date":    a.Date,
				"summary": a.Summary,
			}
		}),
		"certificates": toContext(self.Certificates, func(c Certificate) map[string]any {
			return map[string]any{
				"name":   c.Name,
				"date":   c.Date,
				"issuer": c.Issuer,
				"url":    c.Url,
			}
		}),
		"publications": toContext(self.Publications, func(p Publication) map[string]any {
			return map[string]any{
				"name":        p.Name,
				"publisher":   p.Publisher,
				"releaseDate": p.ReleaseDate,
				"url":         p.Url,
				"summary":     p.Summary,
			}
		}),
		"skills": toContext(self.Skills, func(s Skill) map[string]any {
			return map[string]any{
				"name":     s.Name,
				"level":    s.Level,
				"keywords": s.Keywords,
			}
		}),
		"languages": toContext(self.Languages, func(l Language) map[string]any {
			return map[string]any{
				"language": l.Language,
				"fluency":  l.Fluency,
			}
		}),
		"interests": toContext(self.Interests, func(i Interest) map[string]any {
			return map[string]any{
				"name":     i.Name,
				"keywords": i.Keywords,
			}
		}),
		"references": toContext(self.References, func(r Reference) map[string]any {
			return map[string]any{
				"name":      r.Name,
				"reference": r.Reference,
			}
		}),
		"projects": toContext(self.Projects, func(p Project) map[string]any {
			return map[string]any{
				"name":        p.Name,
				"startDate":   p.StartDate,
				"endDate":     p.EndDate,
				"description": p.Description,
				"highlights":  p.Highlights,
				"url":         p.Url,
			}
		}),
	}
}
