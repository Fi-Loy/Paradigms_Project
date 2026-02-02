// Project CSI2120/CSI2520
// Winter 2026
// Robert Laganiere, uottawa.ca

// this is the (incomplete) Program class
public class Program {
	
	private String programID;
	private String name;
	private int quota;
	private int[] rol;

	Resident[] matchedResidents;

	
	// constructs a Program
    public Program(String id, String n, int q) {
	
		programID= id;
		name= n;
		quota= q;
	}

    // the rol in order of preference
	public void setROL(int[] rol) {
		
		this.rol= rol;
	}
	
	// string representation
	public String toString() {
      
       return "["+programID+"]: "+name+" {"+ quota+ "}" +" ("+rol.length+")";	  
	}

	/**
	 * Checks if resident is matched to this Program
	 * @param residentId
	 */
	public boolean member(int residentId){
		for (Resident resident : matchedResidents) {
			if (resident.getResidentID() == residentId){
				return true;
			}
		}
		return false;
	}

	/**
	 * Returns the rank of the resident within the 
	 * program ROL. Returns -1 if resident is not
	 * included in the ROL. 
	 */
	public int rank(int residentID){
		for(int i = 0; i < rol.length; i++){
			if (rol[i] == residentID)
				return i + 1;
		}
		return -1;
	}

	/**
	 * Returns the resident with the highest rank
	 * (lowest preference)
	 */
	public int leastPreferred(){
		return rol[rol.length];
	}

	/**
	 * Adds Resdient to the match list of the program if
	 * the program has not reached its quota or if the
	 * resident is preferred over an already matched
	 * resident. 
	 */
	public void addResident(Resident resident){
		//TODO: Finish implementation
	}

}