package msg

import (
	"fmt"
	"io"
	"strings"
)

// Pretty writes a human readable representation of Status to w
func (m Status) Pretty(w io.Writer, prefix string) error {
	_, err := fmt.Fprintf(w, "%sStatus:               %s (%d)\n", prefix, m.Message, m.Code)
	if err != nil {
		return err
	}

	if len(m.Location) > 0 {
		_, err = fmt.Fprintf(w, "%sRedirect Location:     %s\n", prefix, m.Location)
	}
	if err != nil {
		return err
	}

	if m.FetchStatus != nil {
		_, err = fmt.Fprintf(w, "%sFetch Status:\n", prefix)
		if err != nil {
			return err
		}
		err = m.FetchStatus.Pretty(w, prefix+"  ")
	}
	if err != nil {
		return err
	}

	return nil
}

// Pretty writes a human readable representation of DataSet to w
func (m DataSet) Pretty(w io.Writer, prefix string) error {
	_, err := fmt.Fprintf(w, "DataSet:\n")
	if err != nil {
		return err
	}

	if m.Categorization != nil {
		err = m.Categorization.Pretty(w, prefix)
		if err != nil {
			return err
		}
	}

	if m.Adfraud != nil {
		err = m.Adfraud.Pretty(w, prefix)
		if err != nil {
			return err
		}
	}

	if m.Malicious != nil {
		err = m.Malicious.Pretty(w, prefix)
		if err != nil {
			return err
		}
	}

	if m.Echo != nil {
		err = m.Echo.Pretty(w, prefix)
		if err != nil {
			return err
		}
	}

	return nil
}

// Pretty writes a human readable representation of Categorization to w
func (m DataSet_Categorization) Pretty(w io.Writer, prefix string) error {
	cats := make([]string, len(m.Category))
	for i, id := range m.Category {
		cats[i] = "\"" + id.Long() + "\""
	}
	_, err := fmt.Fprintf(w, "%sCategorization: %s\n", prefix, strings.Join(cats, " : "))
	return err
}

// Pretty writes a human readable representation of AdFraud to w
func (m DataSet_AdFraud) Pretty(w io.Writer, prefix string) error {
	_, err := fmt.Fprintf(w,
		"%sAd Fraud:\n"+
			"%s  Fraud: %v\n"+
			"%s  Signature: %s\n",
		prefix,
		prefix, m.Fraud,
		prefix, m.Signature,
	)
	return err
}

// Pretty writes a human readable representation of Malicious to w
func (m DataSet_Malicious) Pretty(w io.Writer, prefix string) error {
	_, err := fmt.Fprintf(w,
		"%sMalicious:\n"+
			"%s  Category: %d\n"+
			"%s  Signature: %s\n"+
			"%s  Verdict: %v\n",
		prefix,
		prefix, m.Category,
		prefix, m.Signature,
		prefix, m.Verdict,
	)

	return err
}

// Pretty writes a human readable representation of Echo to w
func (m DataSet_Echo) Pretty(w io.Writer, prefix string) error {
	_, err := fmt.Fprintf(w, "%sEcho: %s\n", prefix, m.Url)
	return err
}

// Pretty writes a human readable representation of QueryResult to w
func (m QueryResult) Pretty(w io.Writer) error {
	_, err := fmt.Fprintf(w,
		"QueryResult:\n"+
			"  URL:                  %s\n"+
			"  Request ID:           %s\n",
		m.Url,
		m.RequestId,
	)
	if err != nil {
		return err
	}

	if len(m.TrackingId) > 0 {
		_, err = fmt.Fprintf(w, "  Tracking ID:          %s\n", m.TrackingId)
		if err != nil {
			return err
		}
	}

	if m.Status != nil {
		err = m.Status.Pretty(w, "  ")
		if err != nil {
			return err
		}
	}

	if len(m.RequestDataset) > 0 {
		ds := make([]string, len(m.RequestDataset))
		for i, d := range m.RequestDataset {
			ds[i] = d.String()
		}

		_, err = fmt.Fprintf(w, "  Requested DataSet(s): %s\n", strings.Join(ds, ", "))
		if err != nil {
			return err
		}
	}

	if m.ResponseDataset != nil {
		_, err = fmt.Fprintf(w, "  Response ")
		if err != nil {
			return err
		}

		err = m.ResponseDataset.Pretty(w, "    ")
		if err != nil {
			return err
		}
	}

	return nil
}
